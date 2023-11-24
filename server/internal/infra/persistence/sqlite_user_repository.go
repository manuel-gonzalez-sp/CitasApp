package persistence

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository(db *gorm.DB) repository.UserRepository {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		utils.Logger.Fatalf("failed to run migration: %v\n", err)
	}
	return &sqliteUserRepository{
		db: db,
	}
}

func (repo *sqliteUserRepository) Create(ctx context.Context, user *entity.User) error {
	result := repo.db.Create(&user)
	return result.Error
}

func (repo *sqliteUserRepository) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	result := repo.db.Delete(&entity.User{}, ID)
	return result.Error
}

func (repo *sqliteUserRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *sqliteUserRepository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.User, error) {
	var user *entity.User
	result := repo.db.First(&user, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *sqliteUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user *entity.User
	result := repo.db.First(&user, "Username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *sqliteUserRepository) UpdateByID(ctx context.Context, ID uuid.UUID, user *entity.User) (*entity.User, error) {
	user.ID = ID
	result := repo.db.Model(user).Updates(*user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
