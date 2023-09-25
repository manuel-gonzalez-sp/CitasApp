package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}

// TODO: Move logic to a new file
// package entity

// import (
// 	"citasapp/internal/infra/utils"

// 	"github.com/google/uuid"
// )

// type User struct {
// 	ID        uuid.UUID
// 	FirstName string
// 	LastName  string
// }

// func NewUser(firstName, lastName string, params ...utils.Param[User]) *User {
// 	user := utils.ApplyParams(&User{
// 		ID:        uuid.New(),
// 		FirstName: firstName,
// 		LastName:  lastName,
// 	}, params)

// 	return user
// }

// func WithID(ID uuid.UUID) utils.Param[User] {
// 	return func(user *User) {
// 		user.ID = ID
// 	}
// }
