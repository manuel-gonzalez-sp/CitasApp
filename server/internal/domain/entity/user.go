package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}

// TODO: WithID optional param
func NewUser(firstName, lastName string) *User {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
	}
}
