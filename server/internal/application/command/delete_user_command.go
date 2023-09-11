package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type DeleteUserCommand struct {
	id uuid.UUID `validate:"required"`
}

func NewDeleteUserCommand(id uuid.UUID) (*DeleteUserCommand, error) {
	cmd := &DeleteUserCommand{
		id: id,
	}
	if err := utils.Validate.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *DeleteUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.DeleteByID(ctx, cmd.id.String())
	if err != nil {
		return nil, err
	}
	return user, nil
}
