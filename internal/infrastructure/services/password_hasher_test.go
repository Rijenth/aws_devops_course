package services_test

import (
	"testing"

	"github.com/rijenth/aws_devops_course/internal/infrastructure/services"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHasher(t *testing.T) {
	hasher := services.NewBcryptPasswordHasher(10)

	password := "mypassword"
	hash, err := hasher.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = hasher.ComparePassword(hash, password)
	assert.NoError(t, err)

	err = hasher.ComparePassword(hash, "wrongpassword")
	assert.Error(t, err)
}
