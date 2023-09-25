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
			utils.WriteErrorResponse(w, http.StatusBadRequest)
			return
		}

		users, usersErr := app.ListUsers(ctx, comm)
		if usersErr != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError)
			return
		}

		utils.WriteJSONResponse(w, http.StatusOK, users)
	}
}
