package repository

import (
	"citasapp/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type FindAllMessagesFilters struct {
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
}

type MessageRepository interface {
	Create(ctx context.Context, message *entity.Message) error
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.Message, error)
	FindAll(ctx context.Context, filters FindAllMessagesFilters) ([]*entity.Message, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, message *entity.Message) (*entity.Message, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}
