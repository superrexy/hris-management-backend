package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func TimeFormatValidation(field validator.FieldLevel) bool {
	_, err := time.Parse("15:04", field.Field().String())
	return err == nil
}

func ISO8601DateTimeValidation(field validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, field.Field().String())
	return err == nil
}
