package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rijenth/aws_devops_course/internal/domain"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB() (*sql.DB, error) {
	// Connexion à une base MySQL temporaire pour les tests
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_database?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	// Insérer un utilisateur fictif
	_, err = db.Exec(`INSERT INTO users (id, username, email, password, created_at) VALUES (1, 'testuser', 'test@example.com', 'hashedpassword', NOW())`)
	require.NoError(t, err)

	user, err := repo.GetUserByUsername(context.Background(), "testuser")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestUserRepository_CreateUser(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	user := &domain.User{
		Username:  "newuser",
		Email:     "new@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
	}

	createdUser, err := repo.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "newuser", createdUser.Username)
}

func TestUserRepository_DeleteUser(t *testing.T) {
	db, err := setupTestDB()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	// Insérer un utilisateur fictif
	_, err = db.Exec(`INSERT INTO users (id, username) VALUES (2, 'deleteuser')`)
	require.NoError(t, err)

	err = repo.DeleteUser(context.Background(), 2)
	assert.NoError(t, err)

	user, err := repo.GetUserByUsername(context.Background(), "deleteuser")
	assert.Error(t, err)
	assert.Nil(t, user)
}
