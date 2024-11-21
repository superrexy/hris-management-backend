package user

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// DI
	userRepository := NewUserRepository()
	userService := NewUserService(userRepository)
	userController := NewUserController(userService)

	// Routes
	userGroup := app.Group("/users")
	userGroup.Get("/", userController.GetAllUsers)
}
