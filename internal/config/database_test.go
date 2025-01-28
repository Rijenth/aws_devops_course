package config_test

import (
	"database/sql"
	"testing"

	"github.com/rijenth/aws_devops_course/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectDatabase_Success(t *testing.T) {
	db, err := config.ConnectDatabase("root", "root", "localhost", "3306", "test_database")
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// Vérifie que la connexion est valide
	err = db.Ping()
	assert.NoError(t, err)

	defer db.Close()
}

func TestConnectDatabase_InvalidCredentials(t *testing.T) {
	// Mauvais mot de passe pour la base de données
	_, err := config.ConnectDatabase("root", "wrong_password", "localhost", "3306", "test_database")
	assert.Error(t, err)
}
