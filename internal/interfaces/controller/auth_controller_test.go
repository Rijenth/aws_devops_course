package controller_test

import (
	"context"
	"errors"
	"testing"

	"github.com/rijenth/aws_devops_course/internal/domain"
	pb "github.com/rijenth/aws_devops_course/internal/grpc/auth"
	"github.com/rijenth/aws_devops_course/internal/interfaces/controller"
	"github.com/rijenth/aws_devops_course/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock du usecase utilisateur
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Authenticate(username, password string) (*domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestLogin_Success(t *testing.T) {
	mockUC := new(MockUserUsecase)
	authCtrl := controller.NewAuthController(mockUC)

	req := &pb.LoginRequest{Username: "testuser", Password: "password"}
	mockUC.On("Authenticate", req.Username, req.Password).Return(&domain.User{Username: "testuser"}, nil)

	res, err := authCtrl.Login(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res.Token)
	mockUC.AssertExpectations(t)
}

func TestLogin_Failure(t *testing.T) {
	mockUC := new(MockUserUsecase)
	authCtrl := controller.NewAuthController(mockUC)

	req := &pb.LoginRequest{Username: "testuser", Password: "wrongpassword"}
	mockUC.On("Authenticate", req.Username, req.Password).Return(nil, errors.New("authentification échouée"))

	res, err := authCtrl.Login(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, res)
	mockUC.AssertExpectations(t)
}
