package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type deleteUserCommand struct {
	ID uuid.UUID `validate:"required"`
}

func NewDeleteUserCommand(ID uuid.UUID) (*deleteUserCommand, error) {
	cmd := &deleteUserCommand{
		ID: ID,
	}
	if err := utils.Validator.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *deleteUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.DeleteByID(ctx, cmd.ID.String())
	return user, err
}
