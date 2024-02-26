package dto

import (
	"citasapp/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID           uuid.UUID  `json:"id" validate:"required"`
	Username     string     `json:"username" validate:"required"`
	Fullname     string     `json:"fullname" validate:"required"`
	BirthDate    time.Time  `json:"birthDate"`
	Gender       string     `json:"gender"`
	Introduction string     `json:"introduction"`
	LookingFor   string     `json:"lookingFor"`
	City         string     `json:"city"`
	Country      string     `json:"country"`
	CreatedAt    time.Time  `json:"createdAt"`
	Photos       []PhotoDTO `json:"photos"`
}
type CreateUserDTO struct {
	Username     string    `json:"username" validate:"required"`
	Fullname     string    `json:"fullname" validate:"required"`
	BirthDate    time.Time `json:"birthDate"`
	Gender       string    `json:"gender"`
	Introduction string    `json:"introduction"`
	LookingFor   string    `json:"lookingFor"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	PhotoURL     string    `json:"photoUrl" validate:"url"`
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
		Photos: []entity.Photo{
			{
				ID:     uuid.New(),
				URL:    dto.PhotoURL,
				IsMain: true,
			},
		},
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

func ToUserDTO(user *entity.User) *UserDTO {
	dto := &UserDTO{
		ID:           user.ID,
		Username:     user.Username,
		Fullname:     user.Fullname,
		BirthDate:    user.BirthDate,
		Gender:       user.Gender,
		Introduction: user.Introduction,
		LookingFor:   user.LookingFor,
		City:         user.City,
		Country:      user.Country,
		CreatedAt:    user.CreatedAt,
		Photos:       []PhotoDTO{},
	}
	if user.Photos != nil {
		for _, value := range user.Photos {
			dto.Photos = append(dto.Photos, *ToPhotoDTO(&value))
		}
	}
	return dto
}

func ToUserDTOs(users []*entity.User) []*UserDTO {
	userDTOs := make([]*UserDTO, 0, len(users))
	for _, user := range users {
		userDTO := ToUserDTO(user)
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs
}
