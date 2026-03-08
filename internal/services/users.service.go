package services

import (
	"api/internal/errors"
	"api/internal/models"
	"api/internal/repositories"
)

type UsersService struct {
	usersRepository repositories.UsersRepository
}

func NewUsersService(usersRepository repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (u UsersService) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.usersRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return user, nil
}
