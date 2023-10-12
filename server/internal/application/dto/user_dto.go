package dto

type CreateUserDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type UpdateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
