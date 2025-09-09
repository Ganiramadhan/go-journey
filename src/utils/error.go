package utils

import (
	"go-journey/src/res"

	"github.com/gofiber/fiber/v2"
)

func HandleValidationError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(
		res.ErrorResponse("Validation failed", err),
	)
}

func HandleServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		res.ErrorResponse("Internal Server Error", err),
	)
}
