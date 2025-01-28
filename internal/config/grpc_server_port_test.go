package config_test

import (
	"testing"

	"github.com/rijenth/aws_devops_course/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetGRPCServerPort_DefaultPort(t *testing.T) {
	port := config.GetGRPCServerPort()
	assert.Equal(t, "50051", port) // Port par défaut
}

func TestGetGRPCServerPort_EnvironmentVariable(t *testing.T) {
	t.Setenv("GRPC_PORT", "60061") // Simule une variable d'environnement
	port := config.GetGRPCServerPort()
	assert.Equal(t, "60061", port)
}
