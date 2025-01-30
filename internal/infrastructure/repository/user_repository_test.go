package repository_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rijenth/aws_devops_course/internal/domain"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestUserRepository_CreateAndGetUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	// Test insertion
	user := &domain.User{Username: "testuser", Password: "hashedpassword"}
	err = userRepo.CreateUser(user)
	assert.NoError(t, err)

	// Test récupération
	retrievedUser, err := userRepo.GetUserByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user.Username, retrievedUser.Username)
}

func TestUserRepository_GetUserByUsername_NotFound(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	retrievedUser, err := userRepo.GetUserByUsername("inexistant")
	assert.Error(t, err)
	assert.Nil(t, retrievedUser)
}
