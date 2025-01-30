package config_test

import (
	"testing"

	"github.com/rijenth/aws_devops_course/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadGrpcServerPortConfig(t *testing.T) {
	config, err := config.LoadGrpcServerPortConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.GrpcServerPort)
}
