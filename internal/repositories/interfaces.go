package repositories

import "api/internal/models"

type AuthRepository interface {
	GenerateToken(code string, challenge string) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
	ValidateToken(token string) (*string, error)
}

type UsersRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
	List(opts *UsersListOptions) ([]models.User, int64, error)
}
