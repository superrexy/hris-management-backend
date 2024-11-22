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
	userGroup.Get("/:id", userController.GetUserByID)
	userGroup.Post("/", userController.CreateUser)
	userGroup.Put("/:id", userController.UpdateUser)
	userGroup.Delete("/:id", userController.DeleteUser)
}
