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
	protected := user.Group("/", middleware.Auth())

	// 🔐 Admin-only routes
	admin := protected.Group("/", middleware.RoleMiddleware("admin"))
	admin.Post("/", controller.CreateUser)
	admin.Put("/:id", controller.UpdateUser)
	admin.Delete("/:id", controller.DeleteUser)
}
