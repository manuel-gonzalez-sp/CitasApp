package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type sendMessageCommand struct {
	senderID   uuid.UUID `validate:"required"`
	receiverID uuid.UUID `validate:"required"`
	content    string    `validate:"required"`
}

func NewSendMessageCommand(content string, senderID, receiverID uuid.UUID) (*sendMessageCommand, error) {
	cmd := &sendMessageCommand{
		senderID:   senderID,
		receiverID: receiverID,
		content:    content,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *sendMessageCommand) Handle(ctx context.Context, messageRepo repository.MessageRepository) (*entity.Message, error) {
	message := entity.NewMessage(cmd.content, cmd.senderID, cmd.receiverID)
	if err := messageRepo.Create(ctx, message); err != nil {
		return nil, err
	}
	return message, nil
}
