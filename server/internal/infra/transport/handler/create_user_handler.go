package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"
)

// @Summary	Create a new User
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.CreateUserDTO true "User DTO"
// @Success	200 {object}	dto.UserDTO
// @Router		/user [post]
func CreateUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		input, inputErr := utils.GetRequestBody[dto.CreateUserDTO](r)
		if inputErr != nil {
			utils.WriteError(w, http.StatusBadRequest, inputErr)
			return
		}

		comm, commErr := command.NewCreateUserCommand(input)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		user, userErr := app.CreateUser(ctx, comm)
		if userErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, userErr)
			return
		}

		dto := dto.ToUserDTO(user)
		utils.WriteResponse(w, http.StatusOK, dto)
	}
}
