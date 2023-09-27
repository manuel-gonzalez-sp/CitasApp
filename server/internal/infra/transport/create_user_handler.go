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

type CreateUserResponse struct {
}

func createUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input, inputErr := utils.GetRequestBody[CreateUserRequest](r)
		if inputErr != nil {
			utils.WriteError(w, http.StatusBadRequest, inputErr)
			return
		}

		comm, commErr := command.NewCreateUserCommand(input.FirstName, input.LastName)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		user, userErr := app.CreateUser(ctx, comm)
		if userErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, userErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, user)
	}
}
