package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type listMessagesCommand struct {
	userID uuid.UUID
	// TOOD: filters, sorting, etc.
}

func NewListMessagesCommand(params ...utils.Param[listMessagesCommand]) (*listMessagesCommand, error) {
	cmd := utils.ApplyParams(&listMessagesCommand{}, params)
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func WithUserID(ID uuid.UUID) utils.Param[listMessagesCommand] {
	return func(cmd *listMessagesCommand) {
		cmd.userID = ID
	}
}

func (cmd *listMessagesCommand) Handle(ctx context.Context, messageRepo repository.MessageRepository) ([]*entity.Message, error) {
	if cmd.userID == uuid.Nil {
		user, err := messageRepo.FindAll(ctx)
		return user, err
	}
	user, err := messageRepo.FindByUserID(ctx, cmd.userID)
	return user, err
}
