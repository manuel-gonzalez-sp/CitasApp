package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/multierr"
)

var (
	Validator  = defaultValidator()
	translator ut.Translator
)

func defaultValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())

	english := en.New()
	universal := ut.New(english, english)

	translator, _ = universal.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, translator)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

func ValidateStruct(value interface{}) (err error) {
	validateErr := Validator.Struct(value)
	if validateErr == nil {
		return nil
	}
	translated := validateErr.(validator.ValidationErrors).Translate(translator)
	for _, translatedErr := range translated {
		err = multierr.Append(err, errors.New(translatedErr))
	}
	return err
}
