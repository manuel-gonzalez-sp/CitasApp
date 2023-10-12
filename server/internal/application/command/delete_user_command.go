package command

import (
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type deleteUserCommand struct {
	userID uuid.UUID `validate:"required"`
}

func NewDeleteUserCommand(userID uuid.UUID) (*deleteUserCommand, error) {
	cmd := &deleteUserCommand{
		userID: userID,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *deleteUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) error {
	return userRepo.DeleteByID(ctx, cmd.userID)
}
