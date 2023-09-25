package utils

import "github.com/go-playground/validator/v10"

var Validator = defaultValidator()

func defaultValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate
}
