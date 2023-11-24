package handler

import (
	"citasapp/internal"
	"citasapp/internal/application/command"
	"citasapp/internal/application/dto"
	"citasapp/internal/infra/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// @Summary	Update a new User
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param input body dto.UpdateUserDTO true "Update DTO"
// @Success	200 {object}	dto.UserDTO
// @Router		/user [patch]
func UpdateUserHandler(app *internal.CitasApp) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		ID, parseErr := uuid.Parse(vars["id"])
		if parseErr != nil {
			utils.WriteError(w, http.StatusBadRequest, parseErr)
			return
		}

		input, inputErr := utils.GetRequestBody[dto.UpdateUserDTO](r)
		if inputErr != nil {
			utils.WriteError(w, http.StatusBadRequest, inputErr)
			return
		}

		comm, commErr := command.NewUpdateUserCommand(ID, input)
		if commErr != nil {
			utils.WriteError(w, http.StatusBadRequest, commErr)
			return
		}

		user, userErr := app.UpdateUser(ctx, comm)
		if userErr != nil {
			utils.WriteError(w, http.StatusInternalServerError, userErr)
			return
		}

		dto := dto.ToUserDTO(user)
		utils.WriteResponse(w, http.StatusOK, dto)
	}
}
