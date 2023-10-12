package repository

import (
	"citasapp/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(ctx context.Context, message *entity.Message) error
	FindByID(ctx context.Context, messageID uuid.UUID) (*entity.Message, error)
	FindAll(ctx context.Context) ([]*entity.Message, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Message, error)
	UpdateByID(ctx context.Context, messageID uuid.UUID, message *entity.Message) (*entity.Message, error)
	DeleteByID(ctx context.Context, messageID uuid.UUID) error
}
