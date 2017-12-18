package server

import (
	"os"
	"errors"
)

type (
	Config struct {
		Port string
	}
)

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("$PORT must set")
	}

	c := &Config{
		Port: port,
	}

	return c, nil
}
