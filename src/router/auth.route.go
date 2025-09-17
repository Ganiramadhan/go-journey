package router

import (
	"go-journey/src/controller"
	"go-journey/src/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	// 🔓 Public routes
	auth.Post("/register", controller.Register)

	// Login with rate limiter
	auth.Post("/login",
		limiter.New(limiter.Config{
			Max:        5,               // max 5 login attempts
			Expiration: 1 * time.Minute, // within 1 minute
			KeyGenerator: func(c *fiber.Ctx) string {
				var body struct {
					Username string `json:"username"`
				}
				_ = c.BodyParser(&body)
				if body.Username != "" {
					return body.Username
				}
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"success": false,
					"message": "Too many login attempts for this account. Please try again later.",
				})
			},
		}),
		controller.Login,
	)

	auth.Post("/refresh", controller.Refresh)

	// 🔒 Protected routes
	auth.Use(middleware.Auth())
	auth.Post("/logout", controller.Logout)
}
