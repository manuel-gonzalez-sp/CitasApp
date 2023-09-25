package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/infra/utils"
	"net/http"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func createUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input, err := utils.GetRequestBody[CreateUserRequest](r)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest)
			return
		}

		comm, commErr := command.NewCreateUserCommand(input.FirstName, input.LastName)
		if commErr != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest)
			return
		}

		user, userErr := app.CreateUser(ctx, comm)
		if userErr != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError)
			return
		}

		utils.WriteJSONResponse(w, http.StatusOK, user)
	}
}
