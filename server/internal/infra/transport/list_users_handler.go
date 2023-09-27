package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/infra/utils"
	"net/http"
)

func listUsersHandler(app *internal.CitasApp) http.HandlerFunc {
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
