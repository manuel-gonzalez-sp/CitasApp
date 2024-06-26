package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/entity"
	"citasapp/internal/domain/repository"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

// Commander represents all the functionality enabled in the application
type Commander interface {
	// Authentication commands
	LogIn(ctx context.Context, cmd *logInCommand) (*dto.LoggedInDTO, error)
	SignUp(ctx context.Context, cmd *signUpCommand) (*dto.LoggedInDTO, error)

	// User commands
	CreateUser(ctx context.Context, cmd *createUserCommand) (*entity.User, error)
	GetUser(ctx context.Context, cmd *getUserCommand) (*entity.User, error)
	ListUsers(ctx context.Context, cmd *listUsersCommand) ([]*entity.User, error)
	UpdateUser(ctx context.Context, cmd *updateUserCommand) (*entity.User, error)
	DeleteUser(ctx context.Context, cmd *deleteUserCommand) error

	// Photo commands
	UploadPhoto(ctx context.Context, cmd *uploadPhotoCommand) (*entity.Photo, error)
	//DeletePhoto(ctx context.Context, cmd *uploadPhotoCommand) (*entity.Photo, error)
}

// DefaultCommanderis the default implementation of CommandHandler
type DefaultCommander struct {
	UserRepo repository.UserRepository
	Bucket   *cloudinary.Cloudinary
}

// Authentication Commands
func (handler *DefaultCommander) LogIn(ctx context.Context, cmd *logInCommand) (*dto.LoggedInDTO, error) {
	return cmd.Handle(ctx, handler.UserRepo)
}
func (handler *DefaultCommander) SignUp(ctx context.Context, cmd *signUpCommand) (*dto.LoggedInDTO, error) {
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

// Photo Commands
func (handler *DefaultCommander) UploadPhoto(ctx context.Context, cmd *uploadPhotoCommand) (*entity.Photo, error) {
	return cmd.Handle(ctx, handler.Bucket)
}
