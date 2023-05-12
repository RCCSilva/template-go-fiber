package main

import (
	"log"

	"rccsilva.com/template-go/config"

	_ "rccsilva.com/template-go/docs"
)

// @title Template Go
// @version 1.0
// @description This is a sample swagger for Fiber
// @host localhost:5000
// @BasePath /
func main() {
	conf := config.New()
	err := conf.Load()
	if err != nil {
		log.Panic("unable to load config ", err)
	}

	app := createApp(conf)
	app.Listen(":5000")
}
