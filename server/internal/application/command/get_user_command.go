package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type getUserCommand struct {
	userID uuid.UUID `validate:"required"`
}

func NewGetUserCommand(userID uuid.UUID) (*getUserCommand, error) {
	cmd := &getUserCommand{
		userID: userID,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *getUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.FindByID(ctx, cmd.userID)
	return user, err
}
