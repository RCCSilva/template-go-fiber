FROM golang:1.20.4-alpine3.17 AS builder
WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o go-app ./cmd/web

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/go-app ./
CMD ["./go-app"]