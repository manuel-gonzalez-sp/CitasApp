package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
	"time"

	"github.com/google/uuid"
)

type readMessageCommand struct {
	messageID uuid.UUID
}

func NewReadMessageCommand(messageID uuid.UUID) (*readMessageCommand, error) {
	cmd := &readMessageCommand{
		messageID: messageID,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *readMessageCommand) Handle(ctx context.Context, messageRepo repository.MessageRepository) (*entity.Message, error) {
	message := &entity.Message{
		ReadAt: time.Now(),
	}
	message, err := messageRepo.UpdateByID(ctx, cmd.messageID, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
