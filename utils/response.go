package utils

import "github.com/gofiber/fiber/v2"

type SuccessResponseParams struct {
	Message    string
	StatusCode int
	Data       interface{}
}

type ErrorResponseParams struct {
	Message    string
	StatusCode int
	Detail     interface{}
}

type successResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type errorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

func SuccessResponse(ctx *fiber.Ctx, params SuccessResponseParams) error {
	if params.StatusCode == 0 {
		params.StatusCode = fiber.StatusOK
	}

	if params.Message == "" {
		params.Message = "Successfully processed the request"
	}

	return ctx.Status(params.StatusCode).JSON(successResponse{
		Status:     "success",
		StatusCode: params.StatusCode,
		Message:    params.Message,
		Data:       params.Data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, params ErrorResponseParams) error {
	if params.StatusCode == 0 {
		params.StatusCode = fiber.StatusInternalServerError
	}

	if params.Message == "" {
		params.Message = "An error occurred while processing the request"
	}

	return ctx.Status(params.StatusCode).JSON(errorResponse{
		Status:     "error",
		StatusCode: params.StatusCode,
		Message:    params.Message,
		Error:      params.Detail,
	})
}
