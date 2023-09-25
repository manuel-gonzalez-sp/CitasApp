package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type createUserCommand struct {
	firstName string `validate:"required"`
	lastName  string `validate:"required"`
}

func NewCreateUserCommand(firstName, lastName string) (*createUserCommand, error) {
	cmd := &createUserCommand{
		firstName: firstName,
		lastName:  lastName,
	}
	if err := utils.Validator.Struct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *createUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user, err := userRepo.Create(ctx, &entity.User{
		FirstName: cmd.firstName,
		LastName:  cmd.lastName,
	})
	return user, err
}
