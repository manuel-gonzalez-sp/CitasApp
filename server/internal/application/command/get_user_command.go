package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type GetUserCommand struct {
	id uuid.UUID `validate:"required"`
}

func NewGetUserCommand(id uuid.UUID) (*GetUserCommand, error) {
	cmd := &GetUserCommand{
		id: id,
	}
	if err := utils.Validate.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *GetUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.FindByID(ctx, cmd.id.String())
	if err != nil {
		return nil, err
	}
	return user, nil
}
