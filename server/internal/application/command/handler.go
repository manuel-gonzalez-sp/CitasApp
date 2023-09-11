package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"
)

// CommandHandler represents all the functionality enabled in the application
type CommandHandler interface {
	GetUser(ctx context.Context, cmd *GetUserCommand) (*entity.User, error)
	CreateUser(ctx context.Context, cmd *CreateUserCommand) (*entity.User, error)
	DeleteUser(ctx context.Context, cmd *DeleteUserCommand) (*entity.User, error)
}

// DefaultCommandHandler is the default implementation of CommandHandler
type DefaultCommandHandler struct {
	UserRepo repository.UserRepository
}

func (handler *DefaultCommandHandler) GetUser(ctx context.Context, cmd *GetUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

func (handler *DefaultCommandHandler) CreateUser(ctx context.Context, cmd *CreateUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

func (handler *DefaultCommandHandler) DeleteUser(ctx context.Context, cmd *DeleteUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
