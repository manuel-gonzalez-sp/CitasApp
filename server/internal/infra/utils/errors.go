package utils

import "errors"

var (
	ErrLogInUserNotFound           = errors.New("user not found")
	ErrLogInWrongPassword          = errors.New("wrong password")
	ErrLogInGeneratingToken        = errors.New("could not generate token")
	ErrSignUpHashingPassword       = errors.New("could not hash password")
	ErrRequiredAuthorizationHeader = errors.New("authorization header is required")
	ErrMalformedAutorizationHeader = errors.New("autorization header is malformed")
	ErrInvalidToken                = errors.New("invalid token")
)
