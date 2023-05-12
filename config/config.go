package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.DatabaseURI = os.Getenv("DATABASE_URI")

	return nil
}
