package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/repository"
	"citasapp/internal/infra/utils"
	"context"
)

type signUpCommand struct {
	dto.SignUpDTO
}

func NewSignUpCommand(signUp dto.SignUpDTO) (*signUpCommand, error) {
	cmd := &signUpCommand{
		SignUpDTO: signUp,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *signUpCommand) Handle(ctx context.Context, userRepo repository.UserRepository) (*dto.LoggedInDTO, error) {
	hashed, err := utils.HashPassword(cmd.Password)
	if err != nil {
		return nil, utils.ErrSignUpHashingPassword
	}

	user := cmd.CreateUserDTO.ToUser()
	user.PasswordHash = hashed
	if err := userRepo.Create(ctx, user); err != nil {
		return nil, err
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
