package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/infra/utils"
	"net/http"

	"github.com/google/uuid"
)

func listMessagesHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := uuid.New() // TODO: Replace with current logged in user

		comm, commErr := command.NewListMessagesCommand(command.WithUserID(userID))
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		messages, messagesErr := app.ListMessages(ctx, comm)
		if messagesErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, messagesErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, messages...)
	}
}
