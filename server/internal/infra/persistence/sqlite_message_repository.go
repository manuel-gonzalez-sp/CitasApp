package persistence

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type sqliteMessageRepository struct {
	db *gorm.DB
}

func NewSQLiteMessageRepository(db *gorm.DB) repository.MessageRepository {
	err := db.AutoMigrate(&entity.Message{})
	if err != nil {
		utils.Logger.Fatalf("failed to run migration: %v\n", err)
	}
	return &sqliteMessageRepository{
		db: db,
	}
}

func (repo *sqliteMessageRepository) Create(ctx context.Context, message *entity.Message) error {
	result := repo.db.Create(&message)
	return result.Error
}

func (repo *sqliteMessageRepository) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	result := repo.db.Delete(&entity.Message{}, ID)
	return result.Error
}

func (repo *sqliteMessageRepository) FindAll(ctx context.Context) ([]*entity.Message, error) {
	var messages []*entity.Message
	result := repo.db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func (repo *sqliteMessageRepository) FindByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Message, error) {
	var messages []*entity.Message
	result := repo.db.
		Where("ReceiverID = ?", userID).
		Or("SenderID = ?", userID).
		Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func (repo *sqliteMessageRepository) FindByID(ctx context.Context, ID uuid.UUID) (*entity.Message, error) {
	var message *entity.Message
	result := repo.db.First(&message, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return message, nil
}

func (repo *sqliteMessageRepository) UpdateByID(ctx context.Context, ID uuid.UUID, message *entity.Message) (*entity.Message, error) {
	message.ID = ID
	result := repo.db.Model(message).Updates(*message)
	if result.Error != nil {
		return nil, result.Error
	}
	return message, nil
}
