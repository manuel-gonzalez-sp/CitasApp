package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type getUserCommand struct {
	ID uuid.UUID `validate:"required"`
}

func NewGetUserCommand(ID uuid.UUID) (*getUserCommand, error) {
	cmd := &getUserCommand{
		ID: ID,
	}
	if err := utils.Validator.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *getUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.FindByID(ctx, cmd.ID)
	return user, err
}
