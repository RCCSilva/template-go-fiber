package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"rccsilva.com/template-go/config"
	"rccsilva.com/template-go/database"
	"rccsilva.com/template-go/domain"

	_ "rccsilva.com/template-go/docs"
)

// @title Template Go
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:5000
// @BasePath /
func main() {
	// Config
	conf := config.New()
	err := conf.Load()
	if err != nil {
		log.Panic("unable to load config ", err)
	}

	// Connect to Database
	db := database.New(conf.DatabaseURI)
	err = db.Connect()
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

	server.Listen(":5000")
}
