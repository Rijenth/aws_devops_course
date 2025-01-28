package usecase_test

import (
	"errors"
	"testing"

	"github.com/rijenth/aws_devops_course/internal/domain"
	"github.com/rijenth/aws_devops_course/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock du repository utilisateur
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUC := usecase.NewUserUsecase(mockRepo)

	user := &domain.User{Username: "nouvel_utilisateur", Password: "motdepasse"}
	mockRepo.On("CreateUser", user).Return(nil)

	err := userUC.CreateUser(user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_UsernameExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUC := usecase.NewUserUsecase(mockRepo)

	user := &domain.User{Username: "existant", Password: "motdepasse"}
	mockRepo.On("CreateUser", user).Return(errors.New("nom d'utilisateur déjà pris"))

	err := userUC.CreateUser(user)
	assert.Error(t, err)
	assert.Equal(t, "nom d'utilisateur déjà pris", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUC := usecase.NewUserUsecase(mockRepo)

	mockUser := &domain.User{Username: "john", Password: "hashedpassword"}
	mockRepo.On("GetUserByUsername", "john").Return(mockUser, nil)

	result, err := userUC.GetUserByUsername("john")

	assert.NoError(t, err)
	assert.Equal(t, "john", result.Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUC := usecase.NewUserUsecase(mockRepo)

	mockRepo.On("GetUserByUsername", "unknown").Return(nil, errors.New("user not found"))

	result, err := userUC.GetUserByUsername("unknown")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
