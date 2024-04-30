package command

import (
	"citasapp/internal/application/dto"
	"citasapp/internal/domain/entity"
	"citasapp/internal/infra/utils"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type uploadPhotoCommand struct {
	dto.UploadPhotoDTO
}

func NewUploadPhotoCommand(uploadPhoto dto.UploadPhotoDTO) (*uploadPhotoCommand, error) {
	cmd := &uploadPhotoCommand{
		uploadPhoto,
	}
	if err := utils.ValidateStruct(cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (cmd *uploadPhotoCommand) Handle(ctx context.Context, bucket *cloudinary.Cloudinary) (*entity.Photo, error) {
	res, err := bucket.Upload.Upload(ctx, cmd.File, uploader.UploadParams{
		Folder: cmd.UserID.String(),
	})
	if err != nil {
		return nil, err
	}
	return &entity.Photo{
		ID:     uuid.New(),
		URL:    res.URL,
		UserID: cmd.UserID,
	}, nil
}
