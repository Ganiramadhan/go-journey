package utils

import "github.com/gofiber/fiber/v2"

func HandleValidationError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "Validation failed",
		"error":   err.Error(),
	})
}

func HandleServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "Internal Server Error",
		"error":   err.Error(),
	})
}
