package middleware

import (
	"go-journey/src/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		token, claims, err := utils.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		if typ, ok := claims["type"].(string); !ok || typ != "access" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token type",
			})
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid subject",
			})
		}

		c.Locals("userID", uint(sub))
		return c.Next()
	}
}
