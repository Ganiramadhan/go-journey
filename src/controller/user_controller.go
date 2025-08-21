package controller

import (
	"strconv"

	"go-journey/src/model"
	"go-journey/src/res"
	"go-journey/src/service"
	"go-journey/src/utils"
	"go-journey/src/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary      Get all users
// @Description  Get list of users
// @Tags         users
// @Produce      json
// @Success      200 {object} res.Response{data=[]model.User}
// @Failure      500 {object} res.Response
// @Router       /users [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return utils.HandleServerError(c, err)
	}
	return c.JSON(res.SuccessResponse("Users fetched successfully", users))
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Get user detail by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200 {object} res.Response{data=model.User}
// @Failure      400 {object} res.Response
// @Failure      404 {object} res.Response
// @Failure      500 {object} res.Response
// @Router       /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.ErrorResponse("Invalid user ID", err.Error()))
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).
				JSON(res.ErrorResponse("User not found", err.Error()))
		}
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("User fetched successfully", user))
}

// CreateUser godoc
// @Summary      Create new user
// @Description  Create a new user with name, email and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      validation.CreateUserRequest  true  "User data"
// @Success      201 {object} res.Response{data=model.User}
// @Failure      400 {object} res.Response
// @Failure      500 {object} res.Response
// @Router       /users [post]
func CreateUser(c *fiber.Ctx) error {
	var req validation.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := service.CreateUser(&user); err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.Status(fiber.StatusCreated).
		JSON(res.SuccessResponse("User created successfully", user))
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int                          true  "User ID"
// @Param        user  body      validation.UpdateUserRequest true "User data"
// @Success      200 {object} res.Response{data=model.User}
// @Failure      400 {object} res.Response
// @Failure      500 {object} res.Response
// @Router       /users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.ErrorResponse("Invalid user ID", err.Error()))
	}

	var req validation.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	user := model.User{
		ID:       uint(id),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := service.UpdateUser(&user); err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("User updated successfully", user))
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200 {object} res.Response
// @Failure      400 {object} res.Response
// @Failure      500 {object} res.Response
// @Router       /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.ErrorResponse("Invalid user ID", err.Error()))
	}

	if err := service.DeleteUser(uint(id)); err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("User deleted successfully", nil))
}
