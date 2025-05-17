package container

import (
	"api/internal/handlers"
	"api/internal/repositories"
	"api/internal/services"

	"gorm.io/gorm"
)

type Container struct {
	//HANDLERS
	AuthHandler *handlers.AuthHandler

	//SERVICES
	AuthService services.AuthService

	//REPOSITORIES
	AuthRepository repositories.AuthRepository
}

func NewContainer(db *gorm.DB) *Container {
	// Initialize repositories
	authRepo := repositories.NewAuthRepository(db)

	// Initialize services
	authService := services.NewAuthService(authRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		//HANDLERS
		AuthHandler: authHandler,

		//SERVICES
		AuthService: authService,

		//REPOSITORIES
		AuthRepository: authRepo,
	}
}
