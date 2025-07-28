package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("Error loading .env file, using default config")
	}
	return &Config{
		Auth: AuthConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
