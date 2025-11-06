package services

import (
	"api/internal/dtos"
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
		return nil, err
	}
	return user, nil
}

func (u UsersService) GetUserByID(id string) (*models.User, error) {
	user, err := u.usersRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(errors.NotFound)
	}
	return user, nil
}

func (u UsersService) ListUsers(options *repositories.UsersListOptions) ([]models.User, int64, error) {
	users, total, err := u.usersRepository.List(options)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (u UsersService) CreateUser(userDto *dtos.CreateUserDto) (*models.User, error) {
	duplicated, err := u.usersRepository.FindByEmail(userDto.Email)
	if err != nil {
		return nil, err
	}
	if duplicated != nil {
		return nil, errors.New(errors.Duplicated)
	}

	user := &models.User{
		Email: userDto.Email,
		Name:  userDto.Name,
	}
	if err := u.usersRepository.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
