package config

import (
	"os"
	"rsch/profile_service/helper"

	"github.com/joho/godotenv"
)

func New() *Config {
	err := godotenv.Load()
	helper.PanicIfError(err)

	server := Server{
		Host: os.Getenv("SERVER_HOST"),
		Port: os.Getenv("SERVER_PORT"),
	}

	database := Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name:     os.Getenv("DATABASE_NAME"),
	}

	return &Config{
		Server:   server,
		Database: database,
	}
}
