package user

import (
	"fmt"
	"hris-management/internal/user/dto"
	"hris-management/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var payload dto.StoreUserPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validation := utils.PayloadValidation(payload); validation != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validation,
		})
	}

	createdUser, err := c.userService.CreateUser(payload)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		StatusCode: fiber.StatusCreated,
		Message:    "User created successfully",
		Data:       createdUser,
	})
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		return err
	}

	fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("An error occurred while fetching the users: %v", err))
	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		StatusCode: fiber.StatusOK,
		Data:       users,
	})
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.GetUserByID(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Successfully fetched user",
		StatusCode: fiber.StatusOK,
		Data:       user,
	})
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payload dto.UpdateUserPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validation := utils.PayloadValidation(payload); validation != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validation,
		})
	}

	updatedUser, err := c.userService.UpdateUser(utils.StringToUint(id), payload)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User updated successfully",
		StatusCode: fiber.StatusOK,
		Data:       updatedUser,
	})
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.GetUserByID(utils.StringToUint(id))
	if err != nil {
		return err
	}

	if err := c.userService.DeleteUser(user); err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User deleted successfully",
		StatusCode: fiber.StatusNoContent,
	})
}
