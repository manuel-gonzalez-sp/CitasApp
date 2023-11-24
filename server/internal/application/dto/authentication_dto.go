package dto

import "citasapp/internal/domain/entity"

type LogInDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoggedInDTO struct {
	entity.User
	JWTToken string `json:"jwtToken"`
}

type SignUpDTO struct {
	CreateUserDTO
	Password string `json:"password" validate:"required"`
}
