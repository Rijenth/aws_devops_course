package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/services"
)

func TestHashPassword(t *testing.T) {
	hasher := services.NewPasswordHasher()

	password := "monMotDePasse"
	hash, err := hasher.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	isValid := hasher.CheckPasswordHash(password, hash)
	assert.True(t, isValid)
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	hasher := services.NewPasswordHasher()

	password := "motDePasse"
	wrongHash := "hashInvalide"
	isValid := hasher.CheckPasswordHash(password, wrongHash)
	assert.False(t, isValid)
}
