package utils

import (
	"go-journey/src/res"

	"github.com/gofiber/fiber/v2"
)

// Send a 400 Bad Request for validation errors
func ValidationError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(
		res.ErrorResponse("Validation failed", err),
	)
}

// Send a 500 Internal Server Error for server errors
func InternalError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		res.ErrorResponse("Internal Server Error", err),
	)
}
