package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/infra/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func readMessageHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		ID, parseErr := uuid.Parse(vars["id"])
		if parseErr != nil {
			utils.WriteError(w, http.StatusBadRequest, parseErr)
			return
		}

		comm, commErr := command.NewReadMessageCommand(ID)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		message, messageErr := app.ReadMessage(ctx, comm)
		if messageErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, messageErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, message)
	}
}
