package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type JwtConfig struct {
	SecretKey string
}

func LoadJwtConfig() (*JwtConfig, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	config := &JwtConfig{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	return config, nil
}
