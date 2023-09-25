package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"
)

// Commander represents all the functionality enabled in the application
type Commander interface {
	CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error)
	GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error)
	ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error)
	DeleteUser(ctx context.Context, cmd *deleteUserCommand) (*entity.User, error)
}

// DefaultCommanderis the default implementation of CommandHandler
type DefaultCommander struct {
	UserRepo repository.UserRepository
}

func (handler *DefaultCommander) CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

func (handler *DefaultCommander) GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

func (handler *DefaultCommander) ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}

func (handler *DefaultCommander) DeleteUser(ctx context.Context, cmd *deleteUserCommand) (*entity.User, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
