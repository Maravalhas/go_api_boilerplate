package clients

import (
	"api/internal/models"
)

type OAuthClient interface {
	GenerateToken(code string, challenge string) (*models.Token, error)
	RefreshToken(refreshToken string) (*models.Token, error)
	ValidateToken(token string) (*string, error)
}
