package command

import (
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"
)

// Commander represents all the functionality enabled in the application
type Commander interface {
	// User commands
	CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error)
	GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error)
	ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error)
	UpdateUser(ctx context.Context, cmd *updateUserCommand) (*entity.User, error)
	DeleteUser(ctx context.Context, cmd *deleteUserCommand) error

	// Message commands
	SendMessage(ctx context.Context, cmd *sendMessageCommand) (*entity.Message, error)
	ReadMessage(ctx context.Context, cmd *readMessageCommand) (*entity.Message, error)
	ListMessages(ctx context.Context, cmd *listMessagesCommand) ([]*entity.Message, error)
}

// DefaultCommanderis the default implementation of CommandHandler
type DefaultCommander struct {
	UserRepo    repository.UserRepository
	MessageRepo repository.MessageRepository
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

// Message Commands
func (handler *DefaultCommander) SendMessage(ctx context.Context, cmd *sendMessageCommand) (*entity.Message, error) {
	return cmd.Handle(ctx, handler.MessageRepo)
}
func (handler *DefaultCommander) ReadMessage(ctx context.Context, cmd *readMessageCommand) (*entity.Message, error) {
	return cmd.Handle(ctx, handler.MessageRepo)
}
func (handler *DefaultCommander) ListMessages(ctx context.Context, cmd *listMessagesCommand) ([]*entity.Message, error) {
	return cmd.Handle(ctx, handler.MessageRepo)
}
