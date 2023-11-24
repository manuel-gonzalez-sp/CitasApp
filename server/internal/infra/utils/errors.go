package utils

import "errors"

var (
	ErrLogInUserNotFound    = errors.New("user not found")
	ErrLogInWrongPassword   = errors.New("wrong password")
	ErrLogInGeneratingToken = errors.New("could not generate token")
)
