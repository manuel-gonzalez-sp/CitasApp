package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type logInCommand struct {
	dto.LogInDTO
}

func NewLogInCommand(logIn dto.LogInDTO) (cmd *logInCommand, err error) {
	cmd.LogInDTO = logIn
	if err = utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *logInCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*dto.LoggedInDTO, error) {
	user, err := userRepo.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return nil, utils.ErrLogInUserNotFound
	}
	if !utils.CheckPasswordHash(cmd.Password, string(user.PasswordHash)) {
		return nil, utils.ErrLogInWrongPassword
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return nil, utils.ErrLogInGeneratingToken

	}
	loggedIn := &dto.LoggedInDTO{
		User:     *user,
		JWTToken: token,
	}
	return loggedIn, err
}
