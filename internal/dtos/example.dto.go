package dtos

import (
	"api/internal/models"
)

type ExampleDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateUserDTO struct {
	Name string `json:"name"`
}

func ToExampleDTO(entity models.Example) ExampleDTO {
	dto := ExampleDTO{
		ID:   entity.ID,
		Name: entity.Name,
	}
	return dto
}

func ToExampleDTOs(entities []models.Example) []ExampleDTO {
	dtos := make([]ExampleDTO, 0, len(entities))
	for _, entity := range entities {
		dtos = append(dtos, ToExampleDTO(entity))
	}
	return dtos
}
