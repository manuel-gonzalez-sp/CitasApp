package repository

import (
	"citasapp/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, ID uuid.UUID) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, user *entity.User) (*entity.User, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error

	// Load from json
}
