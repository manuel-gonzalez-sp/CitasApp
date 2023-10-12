package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"
)

func sendMessageHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input, inputErr := utils.GetRequestBody[dto.SendMessageDTO](r)
		if inputErr != nil {
			utils.WriteError(w, http.StatusBadRequest, inputErr)
			return
		}

		comm, commErr := command.NewSendMessageCommand(input.Content, input.SenderID, input.ReceiverID)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		message, messageErr := app.SendMessage(ctx, comm)
		if messageErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, messageErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, message)
	}
}
