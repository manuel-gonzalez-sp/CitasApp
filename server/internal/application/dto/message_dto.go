package dto

import "github.com/google/uuid"

type SendMessageDTO struct {
	SenderID   uuid.UUID `json:"sender" validate:"required"`
	ReceiverID uuid.UUID `json:"receiver" validate:"required"`
	Content    string    `json:"content" validate:"required"`
}
