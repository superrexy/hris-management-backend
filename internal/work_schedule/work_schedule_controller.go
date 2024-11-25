package workschedule

import (
	"fmt"
	"hris-management/internal/work_schedule/dto"
	"hris-management/utils"

	"github.com/gofiber/fiber/v2"
)

type WorkScheduleController struct {
	workScheduleService WorkScheduleService
}

func NewWorkScheduleController(workScheduleService WorkScheduleService) *WorkScheduleController {
	return &WorkScheduleController{workScheduleService}
}

func (c *WorkScheduleController) CreateWorkSchedule(ctx *fiber.Ctx) error {
	var payload dto.StoreWorkScheduleRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	createdWorkSchedule, err := c.workScheduleService.CreateWorkSchedule(payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Schedule created successfully",
		StatusCode: fiber.StatusCreated,
		Data:       createdWorkSchedule,
	})
}

func (c *WorkScheduleController) DeleteWorkSchedule(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.workScheduleService.DeleteWorkSchedule(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Schedule deleted successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (c *WorkScheduleController) GetAllWorkSchedule(ctx *fiber.Ctx) error {
	workSchedules, err := c.workScheduleService.GetAllWorkSchedule()
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Schedules fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       workSchedules,
	})
}

func (c *WorkScheduleController) GetWorkScheduleByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	workSchedule, err := c.workScheduleService.GetWorkScheduleByID(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Schedule fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       workSchedule,
	})
}

func (c *WorkScheduleController) UpdateWorkSchedule(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payload dto.UpdateWorkScheduleRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	updatedWorkSchedule, err := c.workScheduleService.UpdateWorkSchedule(payload, utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Schedule updated successfully",
		StatusCode: fiber.StatusOK,
		Data:       updatedWorkSchedule,
	})
}

type WorkDayController struct {
	workDayService WorkDayService
}

func NewWorkDayController(workDayService WorkDayService) *WorkDayController {
	return &WorkDayController{
		workDayService: workDayService,
	}
}

func (c *WorkDayController) CreateWorkDay(ctx *fiber.Ctx) error {
	var payload dto.StoreWorkDayRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	createdWorkDay, err := c.workDayService.CreateWorkDay(payload)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Day created successfully",
		StatusCode: fiber.StatusCreated,
		Data:       createdWorkDay,
	})
}

func (c *WorkDayController) DeleteWorkDay(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.workDayService.DeleteWorkDay(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Day deleted successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (c *WorkDayController) GetAllWorkDay(ctx *fiber.Ctx) error {
	workDays, err := c.workDayService.GetAllWorkDay()
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Days fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       workDays,
	})
}

func (c *WorkDayController) GetWorkDayByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	workDay, err := c.workDayService.GetWorkDayByID(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Day fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       workDay,
	})
}

func (c *WorkDayController) UpdateWorkDay(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payload dto.UpdateWorkDayRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	updatedWorkDay, err := c.workDayService.UpdateWorkDay(payload, utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "Work Day updated successfully",
		StatusCode: fiber.StatusOK,
		Data:       updatedWorkDay,
	})
}

type UserWorkScheduleController struct {
	userWorkScheduleService UserWorkScheduleService
}

func NewUserWorkScheduleController(userWorkScheduleService UserWorkScheduleService) *UserWorkScheduleController {
	return &UserWorkScheduleController{
		userWorkScheduleService,
	}
}

func (c *UserWorkScheduleController) CreateUserWorkSchedule(ctx *fiber.Ctx) error {
	var payload dto.StoreUserWorkScheduleRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	createdUserWorkSchedule, err := c.userWorkScheduleService.CreateUserWorkSchedule(payload)
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User Work Schedule created successfully",
		StatusCode: fiber.StatusCreated,
		Data:       createdUserWorkSchedule,
	})
}

func (c *UserWorkScheduleController) DeleteUserWorkSchedule(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.userWorkScheduleService.DeleteUserWorkSchedule(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User Work Schedule deleted successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (c *UserWorkScheduleController) GetAllUserWorkSchedule(ctx *fiber.Ctx) error {
	userWorkSchedules, err := c.userWorkScheduleService.GetAllUserWorkSchedule()
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User Work Schedules fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       userWorkSchedules,
	})
}

func (c *UserWorkScheduleController) GetUserWorkScheduleByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userWorkSchedule, err := c.userWorkScheduleService.GetUserWorkScheduleByID(utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User Work Schedule fetched successfully",
		StatusCode: fiber.StatusOK,
		Data:       userWorkSchedule,
	})
}

func (c *UserWorkScheduleController) UpdateUserWorkSchedule(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payload dto.UpdateUserWorkScheduleRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}

	if validated := utils.PayloadValidation(payload); validated != nil {
		return utils.ErrorResponse(ctx, utils.ErrorResponseParams{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid payload",
			Detail:     validated,
		})
	}

	updatedUserWorkSchedule, err := c.userWorkScheduleService.UpdateUserWorkSchedule(payload, utils.StringToUint(id))
	if err != nil {
		return err
	}

	return utils.SuccessResponse(ctx, utils.SuccessResponseParams{
		Message:    "User Work Schedule updated successfully",
		StatusCode: fiber.StatusOK,
		Data:       updatedUserWorkSchedule,
	})
}
