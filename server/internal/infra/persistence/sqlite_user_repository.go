package persistence

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"

	"github.com/upper/db/v4"
)

type SQLiteUserRepository struct {
	db db.Session
}

func NewSQLiteUserRepository(db db.Session) repository.UserRepository {
	return &SQLiteUserRepository{
		db: db,
	}
}

func (repo *SQLiteUserRepository) Create(ctx context.Context, entity *entity.User) (*entity.User, error) {
	collection := repo.db.Collection("users")
	_, err := collection.Insert(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
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
