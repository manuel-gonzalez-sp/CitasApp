package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type listMessagesCommand struct {
	senderID   uuid.UUID
	receiverID uuid.UUID
	// TOOD: filters, sorting, etc.
}

func NewListMessagesCommand(params ...utils.Param[listMessagesCommand]) (*listMessagesCommand, error) {
	cmd := utils.ApplyParams(&listMessagesCommand{}, params)
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func WithSenderID(ID uuid.UUID) utils.Param[listMessagesCommand] {
	return func(cmd *listMessagesCommand) {
		cmd.senderID = ID
	}
}

func WithReceiverID(ID uuid.UUID) utils.Param[listMessagesCommand] {
	return func(cmd *listMessagesCommand) {
		cmd.receiverID = ID
	}
}

func (cmd *listMessagesCommand) Handle(ctx context.Context, messageRepo repository.MessageRepository) ([]*entity.Message, error) {
	filters := repository.FindAllMessagesFilters{
		SenderID:   cmd.senderID,
		ReceiverID: cmd.receiverID,
	}
	user, err := messageRepo.FindAll(ctx, filters)
	return user, err
}
