package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// @Summary	Get a single User based on the provided ID
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param        id   path    string  true  "User ID"
// @Success	200	{object} dto.UserDTO
// @Router		/user/{id} [get]
func GetUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		ID, parseErr := uuid.Parse(vars["id"])
		if parseErr != nil {
			utils.WriteError(w, http.StatusBadRequest, parseErr)
			return
		}

		comm, commErr := command.NewGetUserCommand(ID)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		user, userErr := app.GetUser(ctx, comm)
		if userErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, userErr)
			return
		}

		dto := dto.ToUserDTO(user)
		utils.WriteResponse(w, http.StatusOK, dto)
	}
}
