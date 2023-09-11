package utils

import "github.com/go-playground/validator/v10"

var Validate = defaultValidator()

func defaultValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}
