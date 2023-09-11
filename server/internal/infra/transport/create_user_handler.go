package transport

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateUserResponse struct {
	dto.UserDTO
}

func createUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		comm, err := command.NewCreateUserCommand("yomero", "kwamero")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		user, err := app.Command.CreateUser(ctx, comm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		out, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Add("content-type", "application/json")
		w.Write(out)

	}
}
