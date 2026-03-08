package services

import (
	"api/internal/dtos"
	"api/internal/errors"
	"api/internal/repositories"
)

type AuthService struct {
	authRepository  repositories.AuthRepository
	usersRepository repositories.UsersRepository
}

func NewAuthService(authRepository repositories.AuthRepository, usersRepository repositories.UsersRepository) *AuthService {
	return &AuthService{
		authRepository:  authRepository,
		usersRepository: usersRepository,
	}
}

func (s AuthService) GenerateToken(code string, challenge string) (*dtos.TokenDTO, error) {
	resp, err := s.authRepository.GenerateToken(code, challenge)
	if err != nil {
		return nil, errors.ErrUnauthorized
	}
	return dtos.ToTokenDTO(resp), nil
}

func (s AuthService) RefreshToken(refreshToken string) (*dtos.TokenDTO, error) {
	resp, err := s.authRepository.RefreshToken(refreshToken)
	if err != nil {
		return nil, errors.ErrUnauthorized
	}
	return dtos.ToTokenDTO(resp), nil
}

func (s AuthService) ValidateToken(token string) (*dtos.UserDto, error) {
	resp, err := s.authRepository.ValidateToken(token)
	if err != nil {
		return nil, errors.ErrUnauthorized
	}
	user, err := s.usersRepository.FindByEmail(*resp)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return dtos.ToUserDto(user), nil
}
