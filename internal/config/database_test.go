package config_test

import (
	"testing"

	"github.com/rijenth/aws_devops_course/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadDatabaseConfig(t *testing.T) {
	config, err := config.LoadDatabaseConfig()
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.DBUser)
	assert.NotEmpty(t, config.DBPassword)
	assert.NotEmpty(t, config.DBName)
	assert.NotEmpty(t, config.DBHost)
	assert.NotEmpty(t, config.DBPort)
}
