package user

import (
	"fmt"
	"hris-management/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{userService}
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
