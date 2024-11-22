package main

import (
	"fmt"
	"hris-management/config"
	userRoute "hris-management/internal/user"
	workScheduleRoute "hris-management/internal/work_schedule"
	"hris-management/utils/exception"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const TAG = "Main::> "

func main() {
	app := fiber.New()

	// Initialize the environment variables
	config.InitENV()

	// Initialize the database connection
	config.InitDB()

	// Middleware
	app.Use(exception.CustomErrorHandler)
	app.Use(recover.New())

	// Routes
	userRoute.SetupRoutes(app)
	workScheduleRoute.SetupRoutes(app)

	// Not Found
	app.Use(func(c *fiber.Ctx) error {
		return exception.NewServiceError(fiber.StatusNotFound, "Not Found", nil)
	})

	address := os.Getenv("APP_ADDRESS")

	fmt.Println(TAG, "Server started at", "http://"+address)
	app.Listen(address)
}
