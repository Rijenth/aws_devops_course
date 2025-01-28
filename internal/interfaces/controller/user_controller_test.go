package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/rijenth/aws_devops_course/internal/domain"
	pb "github.com/rijenth/aws_devops_course/internal/grpc/user"
	"github.com/rijenth/aws_devops_course/internal/interfaces/controller"
	"github.com/rijenth/aws_devops_course/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock du usecase utilisateur
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestGetUser_Success(t *testing.T) {
	mockUC := new(MockUserUsecase)
	userCtrl := controller.NewUserController(mockUC)

	req := &pb.GetUserRequest{Username: "testuser"}
	mockUC.On("GetUserByUsername", req.Username).Return(&domain.User{Username: "testuser", Password: "hashedpassword"}, nil)

	res, err := userCtrl.GetUser(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "testuser", res.Username)
	mockUC.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	mockUC := new(MockUserUsecase)
	userCtrl := controller.NewUserController(mockUC)

	req := &pb.GetUserRequest{Username: "unknown"}
	mockUC.On("GetUserByUsername", req.Username).Return(nil, errors.New("user not found"))

	res, err := userCtrl.GetUser(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, res)
	mockUC.AssertExpectations(t)
}
