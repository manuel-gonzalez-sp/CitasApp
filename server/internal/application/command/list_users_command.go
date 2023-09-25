package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type listUsersCommand struct {
	// TOOD: filters, sorting, etc.
}

func NewListUsersCommand() (*listUsersCommand, error) {
	cmd := &listUsersCommand{}
	if err := utils.Validator.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *listUsersCommand) Handle(ctx context.Context, userRepo repository.UserRepository) ([]*entity.User, error) {
	user, err := userRepo.FindAll(ctx)
	return user, err
}
