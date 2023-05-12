package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"rccsilva.com/template-go/config"
	"rccsilva.com/template-go/domain"
)

func TestCreateUser(t *testing.T) {
	conf := &config.Config{DatabaseURI: "postgres://user:Password!23@localhost:5432/template-go?sslmode=disable"}
	app := createApp(conf)

	makeRequest := func(payload any) (*http.Response, error) {
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(fiber.MethodPost, "/api/v1/user", bytes.NewReader(body))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return app.Test(req)
	}

	t.Run("it creates an user given email, username and password", func(t *testing.T) {
		// Arrange
		payload := domain.CreateUserRequest{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		// Act
		res, _ := makeRequest(payload)

		// Assert
		assert.Equal(t, fiber.StatusCreated, res.StatusCode)
	})
}
