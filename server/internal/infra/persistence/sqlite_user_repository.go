package persistence

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"
)

type SQLiteUserRepository struct{}

// Create implements repository.UserRepository.
func (*SQLiteUserRepository) Create(ctx context.Context, entity *entity.User) (*entity.User, error) {
	panic("unimplemented")
}

// DeleteByID implements repository.UserRepository.
func (*SQLiteUserRepository) DeleteByID(ctx context.Context, id string) (*entity.User, error) {
	panic("unimplemented")
}

// FindAll implements repository.UserRepository.
func (*SQLiteUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	panic("unimplemented")
}

// FindByID implements repository.UserRepository.
func (*SQLiteUserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	panic("unimplemented")
}

// UpdateByID implements repository.UserRepository.
func (*SQLiteUserRepository) UpdateByID(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	panic("unimplemented")
}

func NewSQLiteUserRepository() repository.UserRepository {
	return &SQLiteUserRepository{}
}
