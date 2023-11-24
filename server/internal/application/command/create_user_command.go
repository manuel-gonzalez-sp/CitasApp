package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type createUserCommand struct {
	dto.CreateUserDTO
}

func NewCreateUserCommand(createUser dto.CreateUserDTO) (*createUserCommand, error) {
	cmd := &createUserCommand{
		CreateUserDTO: createUser,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *createUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user := cmd.CreateUserDTO.ToUser()
	if err := userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
