package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type CreateUserCommand struct {
	firstName string `validate:"required"`
	lastName  string `validate:"required"`
}

func NewCreateUserCommand(firstName string, lastName string) (*CreateUserCommand, error) {
	cmd := &CreateUserCommand{
		firstName: firstName,
		lastName:  lastName,
	}
	if err := utils.Validate.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *CreateUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.Create(ctx, entity.NewUser(
		cmd.firstName,
		cmd.lastName,
	))
	if err != nil {
		return nil, err
	}
	return user, nil
}
