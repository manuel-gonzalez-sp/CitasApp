package dto

import (
	"citasapp/internal/domain/entity"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func (dto *UserDTO) ToUser() entity.User {
	return entity.User{
		ID:        dto.ID,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}

func UserToDTO(user *entity.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

// TODO: from/to JSON
