package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func TimeFormatValidation(field validator.FieldLevel) bool {
	_, err := time.Parse("15:04", field.Field().String())
	return err == nil
}
