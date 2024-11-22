package exception

import (
	"hris-management/utils"

	"github.com/gofiber/fiber/v2"
)

type ServiceError struct {
	StatusCode int
	Message    string
	Detail     interface{}
}

func (s *ServiceError) Error() string {
	return s.Message
}

func NewServiceError(statusCode int, message string, detail interface{}) *ServiceError {
	return &ServiceError{
		StatusCode: statusCode,
		Message:    message,
		Detail:     detail,
	}
}

func CustomErrorHandler(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err != nil {
		if customErr, ok := err.(*ServiceError); ok {
			return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
				StatusCode: customErr.StatusCode,
				Message:    customErr.Message,
				Detail:     customErr.Detail,
			})
		}

		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "An error occurred",
			Detail:     err.Error(),
		})
	}

	return nil
}
