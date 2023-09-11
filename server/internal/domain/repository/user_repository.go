package repository

import (
	"citasapp/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, entity *entity.User) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	UpdateByID(ctx context.Context, id string, user *entity.User) (*entity.User, error)
	DeleteByID(ctx context.Context, id string) (*entity.User, error)
}
