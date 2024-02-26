package dto

import (
	"citasapp/internal/domain/entity"

	"github.com/google/uuid"
)

type PhotoDTO struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	URL    string    `json:"url" validate:"url"`
	IsMain bool      `json:"isMain"`
}

func ToPhotoDTO(photo *entity.Photo) *PhotoDTO {
	dto := &PhotoDTO{
		ID:     photo.ID,
		URL:    photo.URL,
		IsMain: photo.IsMain,
	}
	return dto
}
