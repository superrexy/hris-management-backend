package workschedule

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// DI
	workScheduleRepository := NewWorkScheduleRepository()
	workScheduleService := NewWorkScheduleService(workScheduleRepository)
	workScheduleController := NewWorkScheduleController(workScheduleService)

	workDayRepository := NewWorkDayRepository()
	workDayService := NewWorkDayService(workDayRepository, workScheduleRepository)
	workDayController := NewWorkDayController(workDayService)

	// Routes
	workScheduleGroup := app.Group("/work-schedules")
	workScheduleGroup.Get("/", workScheduleController.GetAllWorkSchedule)
	workScheduleGroup.Get("/:id", workScheduleController.GetWorkScheduleByID)
	workScheduleGroup.Post("/", workScheduleController.CreateWorkSchedule)
	workScheduleGroup.Put("/:id", workScheduleController.UpdateWorkSchedule)
	workScheduleGroup.Delete("/:id", workScheduleController.DeleteWorkSchedule)

	workDayGroup := app.Group("/work-days")
	workDayGroup.Get("/", workDayController.GetAllWorkDay)
	workDayGroup.Get("/:id", workDayController.GetWorkDayByID)
	workDayGroup.Post("/", workDayController.CreateWorkDay)
	workDayGroup.Put("/:id", workDayController.UpdateWorkDay)
	workDayGroup.Delete("/:id", workDayController.DeleteWorkDay)
}
