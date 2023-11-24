package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type updateUserCommand struct {
	dto.UpdateUserDTO
	ID uuid.UUID `validate:"required"`
}

func NewUpdateUserCommand(ID uuid.UUID, updateUser dto.UpdateUserDTO) (*updateUserCommand, error) {
	cmd := &updateUserCommand{
		UpdateUserDTO: updateUser,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *updateUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user := cmd.UpdateUserDTO.ToUser()
	user, err := userRepo.UpdateByID(ctx, cmd.ID, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
