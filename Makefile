tidy:
	go mod tidy

swagger:
	swag init -g ./cmd/web/main.go

test:
	go test ./...