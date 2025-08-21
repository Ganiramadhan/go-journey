package handler

import (
	"github.com/ganiramadhan/go-fiber-app/internal/repository"
	"github.com/ganiramadhan/go-fiber-app/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users := h.service.GetUsers()
	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user repository.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	created := h.service.CreateUser(user)
	return c.Status(fiber.StatusCreated).JSON(created)
}
