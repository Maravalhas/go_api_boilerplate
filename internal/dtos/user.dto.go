package dtos

import (
	"api/internal/models"
	"github.com/google/uuid"
)

type UserDto struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

func ToUserDto(user *models.User) *UserDto {
	return &UserDto{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

type CreateUserDto struct {
	Email string `form:"email" json:"email" binding:"required,email"`
	Name  string `form:"name" json:"name" binding:"required"`
}
