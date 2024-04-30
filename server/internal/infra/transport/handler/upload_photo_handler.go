package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"

	"github.com/google/uuid"
)

// @Summary	Uploads a new Photo
// @Security ApiKeyAuth
// @Tags Photos
// @Accept json
// @Produce json
// @Param input body dto.UploadPhotoDTO true "Photo DTO"
// @Success	200 {object}	dto.PhotoDTO
// @Router		/photo [post]
func UploadPhotoHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		r.ParseMultipartForm(10 << 20) // Limit upload size

		file, _, err := r.FormFile("upload")
		if err != nil {
			http.Error(w, "Invalid file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		user, _ := utils.GetUserClaimsFromContext(r.Context())

		comm, commErr := command.NewUploadPhotoCommand(dto.UploadPhotoDTO{
			UserID: uuid.MustParse(user),
			File:   file,
		})
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		photo, photoErr := app.UploadPhoto(ctx, comm)
		if photoErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, photoErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, photo)
	}
}
