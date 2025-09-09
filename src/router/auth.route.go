package router

import (
	"go-journey/src/controller"
	"go-journey/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	// 🔓 Public routes
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Post("/refresh", controller.Refresh)

	// 🔒 Protected routes
	auth.Use(middleware.Auth())
	auth.Post("/logout", controller.Logout)
	auth.Get("/check", controller.CheckToken)
}
