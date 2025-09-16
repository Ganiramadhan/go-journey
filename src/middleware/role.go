package middleware

import (
	"go-journey/src/database"
	"go-journey/src/model"

	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid user id in context",
			})
		}

		var user model.User
		if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "user not found",
			})
		}

		for _, role := range allowedRoles {
			if user.Role == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "access denied for your role",
		})
	}
}
