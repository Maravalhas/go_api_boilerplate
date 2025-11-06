package container

import (
	"api/internal/clients"
	"api/internal/config"
	"api/internal/handlers"
	"api/internal/middlewares"
	"api/internal/repositories"
	"api/internal/services"

	"gorm.io/gorm"
)

type Container struct {
	//MIDDLEWARES
	Middlewares *middlewares.Middlewares

	//HANDLERS
	AuthHandler *handlers.AuthHandler

	//SERVICES
	AuthService  *services.AuthService
	UsersService *services.UsersService

	//REPOSITORIES
	AuthRepository  repositories.AuthRepository
	UsersRepository repositories.UsersRepository

	//CLIENTS
	OAuthClient clients.OAuthClient
}

func NewContainer(db *gorm.DB, config *config.Config) *Container {
	// Initialize clients
	oauthClient := clients.NewOAuthClient(config)

	// Initialize repositories
	authRepo := repositories.NewAuthRepository(db, oauthClient)
	usersRepo := repositories.NewUsersRepository(db)

	// Initialize services
	authService := services.NewAuthService(authRepo, usersRepo)
	usersService := services.NewUsersService(usersRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize middlewares
	Middlewares := middlewares.New(authService, usersService)

	return &Container{
		//MIDDLEWARES
		Middlewares: Middlewares,

		//HANDLERS
		AuthHandler: authHandler,

		//SERVICES
		AuthService:  authService,
		UsersService: usersService,

		//REPOSITORIES
		AuthRepository:  authRepo,
		UsersRepository: usersRepo,

		//CLIENTS
		OAuthClient: oauthClient,
	}
}
