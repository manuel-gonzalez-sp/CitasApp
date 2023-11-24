package dto

import (
	"citasapp/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

type CreateUserDTO struct {
	Username     string    `json:"username" validate:"required"`
	Fullname     string    `json:"fullname" validate:"required"`
	BirthDate    time.Time `json:"birthDate"`
	Gender       string    `json:"gender"`
	Introduction string    `json:"introduction"`
	LookingFor   string    `json:"lookingFor"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
}

type UpdateUserDTO struct {
	Fullname     string    `json:"fullname"`
	BirthDate    time.Time `json:"birthDate"`
	Gender       string    `json:"gender"`
	Introduction string    `json:"introduction"`
	LookingFor   string    `json:"lookingFor"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
}

func (dto *CreateUserDTO) ToUser() *entity.User {
	return &entity.User{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		Username:     dto.Username,
		Fullname:     dto.Fullname,
		BirthDate:    dto.BirthDate,
		Gender:       dto.Gender,
		Introduction: dto.Introduction,
		LookingFor:   dto.LookingFor,
		City:         dto.City,
		Country:      dto.Country,
	}
}

func (dto *UpdateUserDTO) ToUser() *entity.User {
	return &entity.User{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		Fullname:     dto.Fullname,
		BirthDate:    dto.BirthDate,
		Gender:       dto.Gender,
		Introduction: dto.Introduction,
		LookingFor:   dto.LookingFor,
		City:         dto.City,
		Country:      dto.Country,
	}
}
