package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validation = validator.New()

type errorValidation struct {
	Field string      `json:"field"`
	Error string      `json:"error"`
	Param interface{} `json:"param"`
}

func PayloadValidation(s interface{}) (errors []errorValidation) {
	if err := validation.Struct(s); err != nil {
		errors := make([]errorValidation, 0)
		for _, err := range err.(validator.ValidationErrors) {
			if param := err.Param(); param != "" {
				errors = append(errors, errorValidation{
					Field: strings.ToLower(err.Field()),
					Error: err.Tag(),
					Param: param,
				})
				continue
			}

			errors = append(errors, errorValidation{
				Field: strings.ToLower(err.Field()),
				Error: err.Tag(),
			})
		}
		return errors
	}

	return nil
}
