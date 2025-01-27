package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GrpcServerPortConfig struct {
	GrpcServerPort string
}

func LoadGrpcServerPortConfig() (*GrpcServerPortConfig, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	config := &GrpcServerPortConfig{
		GrpcServerPort: os.Getenv("GRPC_SERVER_PORT"),
	}

	return config, nil
}
