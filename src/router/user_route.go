package router

import (
	"go-journey/src/controller"
	"go-journey/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users")

	// 🔓 Public routes
	user.Get("/", controller.GetUsers)
	user.Get("/:id", controller.GetUser)

	// 🔒 Protected routes
	user.Use(middleware.Auth())
	user.Post("/", controller.CreateUser)
	user.Put("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)
}
