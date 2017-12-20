package server

import (
	"os"
	"errors"
)

type (
	Config struct {
		Port       string
		MongodbURI string
	}
)

func GetConfig() (*Config, error) {
	port := os.Getenv("PORT")
	mgoURI := os.Getenv("MONGODB_URI")

	if port == "" {
		return nil, errors.New("$PORT must set")
	}
	if mgoURI == "" {
		return nil, errors.New("$MONGODB_URI must set")
	}

	c := &Config{
		Port:       port,
		MongodbURI: mgoURI,
	}

	return c, nil
}
