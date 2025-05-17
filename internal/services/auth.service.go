package services

import (
	"api/internal/repositories"
	"api/internal/structs"
)

type AuthService interface {
	GenerateToken(code string, challenge string) (*structs.GoogleTokenResponse, error)
	ValidateToken(token string) (*structs.GoogleTokenValidationResponse, error)
	RefreshToken(refreshToken string) (*structs.GoogleTokenRefreshResponse, error)
}

type authService struct {
	AuthRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &authService{
		AuthRepository: authRepository,
	}
}

func (s *authService) GenerateToken(code string, challenge string) (*structs.GoogleTokenResponse, error) {
	resp, err := s.AuthRepository.GenerateToken(code, challenge)
	return resp, err
}

func (s *authService) ValidateToken(token string) (*structs.GoogleTokenValidationResponse, error) {
	resp, err := s.AuthRepository.ValidateToken(token)
	return resp, err
}

func (s *authService) RefreshToken(refreshToken string) (*structs.GoogleTokenRefreshResponse, error) {
	resp, err := s.AuthRepository.RefreshToken(refreshToken)
	return resp, err
}
