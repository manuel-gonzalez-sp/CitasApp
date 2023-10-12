package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/google/uuid"
)

type updateUserCommand struct {
	ID        uuid.UUID `validate:"required"`
	firstName string
	lastName  string
}

func NewUpdateUserCommand(ID uuid.UUID, firstName, lastName string) (*updateUserCommand, error) {
	cmd := &updateUserCommand{
		ID:        ID,
		firstName: firstName,
		lastName:  lastName,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *updateUserCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*entity.User, error) {
	user := entity.NewUser(cmd.firstName, cmd.lastName)
	user, err := userRepo.UpdateByID(ctx, cmd.ID, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
