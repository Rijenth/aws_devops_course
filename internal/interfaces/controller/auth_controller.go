package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rijenth/aws_devops_course/internal/contextkeys"
	"github.com/rijenth/aws_devops_course/internal/domain"
	pbAuth "github.com/rijenth/aws_devops_course/internal/grpc/auth"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/services"
	"github.com/rijenth/aws_devops_course/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	passwordHasher *services.BcryptPasswordHasher
	usecase        *usecase.UserUsecase
	JWTsecretKey   string
	pbAuth.UnimplementedAuthServiceServer
}

func NewAuthController(uc *usecase.UserUsecase, bcryptPasswordHasher *services.BcryptPasswordHasher, secretKey string) *AuthController {
	return &AuthController{usecase: uc, passwordHasher: bcryptPasswordHasher, JWTsecretKey: secretKey}
}

func (authController *AuthController) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	user, err := authController.usecase.GetUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if err := authController.passwordHasher.ComparePassword(user.Password, req.Password); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "authentication failed")
	}

	expiresIn := time.Hour
	expirationTime := time.Now().Add(expiresIn)

	claims := &domain.JWTClaims{
		UserID:   strconv.Itoa(user.ID),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(authController.JWTsecretKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate token: %v", err)
	}

	authController.usecase.UpdateUserLoginState(ctx, user, true)

	return &pbAuth.LoginResponse{
		Token:     tokenString,
		ExpiresIn: int64(expiresIn.Seconds()),
	}, nil
}

func (authController *AuthController) UpdatePassword(ctx context.Context, req *pbAuth.UpdatePasswordRequest) (*pbAuth.UpdatePasswordResponse, error) {
	user, err := authController.usecase.GetUserByUsername(ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if err := authController.passwordHasher.ComparePassword(user.Password, req.OldPassword); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "aborting request due to incorrect password")
	}

	hashedPassword, err := authController.passwordHasher.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not hash password: %v", err)
	}

	user.Password = hashedPassword

	_, err = authController.usecase.UpdateUserPassword(ctx, user)

	if err != nil {
		return nil, err
	}

	return &pbAuth.UpdatePasswordResponse{
		Success: true,
	}, nil
}

func (authController *AuthController) Logout(ctx context.Context, req *pbAuth.LogoutRequest) (*pbAuth.LogoutResponse, error) {
	authenticatedUsername, _ := ctx.Value(contextkeys.AuthenticatedUserUsernameKey).(string)

	if authenticatedUsername == "" {
		return &pbAuth.LogoutResponse{
			Success: true,
		}, nil
	}

	user, _ := authController.usecase.GetUserByUsername(ctx, authenticatedUsername)

	if user != nil {
		authController.usecase.UpdateUserLoginState(ctx, user, false)
	}

	return &pbAuth.LogoutResponse{
		Success: true,
	}, nil
}
