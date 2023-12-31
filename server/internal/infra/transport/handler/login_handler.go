package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"
)

// @Summary User login
// @Description Logs in a user with provided credentials.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body dto.LogInDTO true "Login Credentials"
// @Success 200 {object} dto.LoggedInDTO "Successful login"
// @Router /login [post]
func LogInHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input, inputErr := utils.GetRequestBody[dto.LogInDTO](r)
		if inputErr != nil {
			utils.WriteError(w, http.StatusBadRequest, inputErr)
			return
		}

		comm, commErr := command.NewLogInCommand(input)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		login, loginErr := app.LogIn(ctx, comm)
		if loginErr == utils.ErrLogInUserNotFound || loginErr == utils.ErrLogInWrongPassword {
			utils.WriteError(w, http.StatusUnauthorized, loginErr)
			return
		}
		if loginErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, loginErr)
			return
		}

		utils.WriteResponse(w, http.StatusOK, login)
	}
}
