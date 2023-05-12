package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"rccsilva.com/template-go/config"
	"rccsilva.com/template-go/database"
	"rccsilva.com/template-go/domain"
)

func createApp(conf *config.Config) *fiber.App {
	// Config

	// Connect to Database
	db := database.New(conf.DatabaseURI)
	err := db.Connect()
	if err != nil {
		log.Panic("unable to connect to database", err)
	}

	// // Migrate
	err = db.Migrate()
	if err != nil {
		log.Panic("unable to migrate", err)
	}

	// Start Server
	server := fiber.New()
	app := domain.New(db)
	handlers := newHandler(app)
	handlers.configRoutes(server)

	return server
}
