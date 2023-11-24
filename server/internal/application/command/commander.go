package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"
)

// Commander represents all the functionality enabled in the application
type Commander interface {
	// Authentication commands
	LogIn(ctx context.Context, cmd *logInCommand) (*dto.LoggedInDTO, error)

	// User commands
	CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error)
	GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error)
	ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error)
	UpdateUser(ctx context.Context, cmd *updateUserCommand) (*entity.User, error)
	DeleteUser(ctx context.Context, cmd *deleteUserCommand) error
}

// DefaultCommanderis the default implementation of CommandHandler
type DefaultCommander struct {
	UserRepo repository.UserRepository
}

// Authentication Commands
func (handler *DefaultCommander) LogIn(ctx context.Context, cmd *logInCommand) (*dto.LoggedInDTO, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

// User Commands
func (handler *DefaultCommander) CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
func (handler *DefaultCommander) GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
func (handler *DefaultCommander) ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
func (handler *DefaultCommander) UpdateUser(ctx context.Context, cmd *updateUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
func (handler *DefaultCommander) DeleteUser(ctx context.Context, cmd *deleteUserCommand) error {
	return cmd.Handle(ctx, handler.UserRepo)
}
