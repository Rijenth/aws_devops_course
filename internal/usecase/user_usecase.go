package usecase

import (
	"context"
	"time"

	"github.com/rijenth/gRPC/internal/contextkeys"
	"github.com/rijenth/gRPC/internal/domain"
	pbUser "github.com/rijenth/gRPC/internal/grpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUserLoginState(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserUsecase struct {
	repository UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{repository: repository}
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	return uc.repository.GetAllUsers(ctx)
}

func (uc *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return uc.repository.GetUserByUsername(ctx, username)
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return uc.repository.CreateUser(ctx, user)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, request *pbUser.UpdateUserRequest) (*domain.User, error) {
	user, err := uc.repository.GetUserByID(ctx, int(request.Id))

	if err != nil {
		return nil, err
	}

	if request.Username != "" {
		authenticatedUserUsername := ctx.Value(contextkeys.AuthenticatedUserUsernameKey).(string)

		existingUser, err := uc.repository.GetUserByUsername(ctx, request.Username)

		if err != nil {
			code := status.Code(err)

			if code != codes.NotFound {
				return nil, err
			}
		}

		if existingUser != nil && existingUser.ID != int(request.Id) {
			return nil, status.Errorf(codes.AlreadyExists, "This username is already taken")
		}

		if request.Username != authenticatedUserUsername {
			user.Username = request.Username
		}
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.FirstName != "" {
		user.FirstName = request.FirstName
	}

	if request.LastName != "" {
		user.LastName = request.LastName
	}

	if request.DateOfBirth != nil {
		user.DateOfBirth = request.DateOfBirth.AsTime()
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	if request.PhoneNumber != "" {
		user.PhoneNumber = request.PhoneNumber
	}

	if request.ProfilePicture != "" {
		user.ProfilePicture = request.ProfilePicture
	}

	if request.Bio != "" {
		user.Bio = request.Bio
	}

	if request.IsActive {
		user.IsActive = request.IsActive
	}

	if request.IsAdmin {
		user.IsAdmin = request.IsAdmin
	}

	return uc.repository.UpdateUser(ctx, user)
}

func (uc *UserUsecase) UpdateUserPassword(ctx context.Context, user *domain.User) (*domain.User, error) {
	return uc.repository.UpdateUserPassword(ctx, user)
}

func (uc *UserUsecase) UpdateUserLoginState(ctx context.Context, user *domain.User, isActive bool) (*domain.User, error) {
	if isActive {
		currentTime := time.Now()

		user.LastLogin = &currentTime
	}

	user.IsActive = isActive

	return uc.repository.UpdateUserLoginState(ctx, user)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return uc.repository.DeleteUser(ctx, id)
}
