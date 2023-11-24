package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/infra/utils"
	"net/http"
)

// @Summary	Get all users
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Success	200	{array}	entity.User
// @Router		/user [get]
func ListUsersHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		comm, commErr := command.NewListUsersCommand()
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		users, usersErr := app.ListUsers(ctx, comm)
		if usersErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, usersErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, users...)
	}
}
