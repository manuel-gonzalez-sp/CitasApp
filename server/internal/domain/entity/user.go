package entity

import (
	"citasapp/internal/infra/utils"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func NewUser(firstName, lastName string, params ...utils.Param[User]) *User {
	user := utils.ApplyParams(&User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
	}, params)
	return user
}

func WithID(ID uuid.UUID) utils.Param[User] {
	return func(user *User) {
		user.ID = ID
	}
}
