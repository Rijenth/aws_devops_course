package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/rijenth/aws_devops_course/internal/config"
	pbauth "github.com/rijenth/aws_devops_course/internal/grpc/auth"
	pbuser "github.com/rijenth/aws_devops_course/internal/grpc/user"
	infrastructure "github.com/rijenth/aws_devops_course/internal/infrastructure/database"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/interceptors"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/repository"
	"github.com/rijenth/aws_devops_course/internal/infrastructure/services"
	"github.com/rijenth/aws_devops_course/internal/interfaces/controller"
	"github.com/rijenth/aws_devops_course/internal/usecase"

	"google.golang.org/grpc"
)

func main() {
	db, err := infrastructure.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize the database: %v", err)
	}
	defer db.Close()

	grpcServerPortConfig, err := config.LoadGrpcServerPortConfig()
	if err != nil {
		log.Fatalf("failed to load gRPC server port config: %v", err)
	}

	jwtSecretKeyConfig, err := config.LoadJwtConfig()
	if err != nil {
		log.Fatalf("failed to load JWT secret")
	}

	authController := setupAuthController(db, jwtSecretKeyConfig)
	userController := setupUserController(db)

	lis, err := net.Listen("tcp", ":"+grpcServerPortConfig.GrpcServerPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.UnaryServerInterceptor(jwtSecretKeyConfig.SecretKey)),
	)
	pbauth.RegisterAuthServiceServer(grpcServer, authController)
	pbuser.RegisterUserServiceServer(grpcServer, userController)

	// TODO: créer une variable d'environnement pour ne pas
	// activer cette fonctionnalité lors d'une éventuelle mise en production
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func setupUserController(db *sql.DB) *controller.UserController {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	return controller.NewUserController(userUsecase)
}

func setupAuthController(db *sql.DB, jwtSecretKeyConfig *config.JwtConfig) *controller.AuthController {
	passwordHasher := services.NewBcryptPasswordHasher(10)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	return controller.NewAuthController(userUsecase, passwordHasher, jwtSecretKeyConfig.SecretKey)
}
