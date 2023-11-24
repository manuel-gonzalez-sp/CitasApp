package dto

type LogInDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoggedInDTO struct {
	UserDTO
	JWTToken string `json:"jwtToken"`
}

type SignUpDTO struct {
	CreateUserDTO
	Password string `json:"password" validate:"required"`
}
