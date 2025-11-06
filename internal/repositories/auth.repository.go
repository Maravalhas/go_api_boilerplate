package repositories

import (
	"api/internal/clients"
	"api/internal/models"

	"gorm.io/gorm"
)

type authRepository struct {
	DB          *gorm.DB
	oauthClient clients.OAuthClient
}

func NewAuthRepository(db *gorm.DB, oauthClient clients.OAuthClient) AuthRepository {
	return authRepository{
		DB:          db,
		oauthClient: oauthClient,
	}
}

func (r authRepository) GenerateToken(code string, challenge string) (*models.Token, error) {
	return r.oauthClient.GenerateToken(code, challenge)
}

func (r authRepository) RefreshToken(refreshToken string) (*models.Token, error) {
	return r.oauthClient.RefreshToken(refreshToken)
}

func (r authRepository) ValidateToken(token string) (*string, error) {
	return r.oauthClient.ValidateToken(token)
}
