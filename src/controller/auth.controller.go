package controller

import (
	"strconv"

	"go-journey/src/database"
	"go-journey/src/model"
	"go-journey/src/res"
	"go-journey/src/utils"
	"go-journey/src/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ===================== REGISTER =====================
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body RegisterRequest true "Register payload"
// @Success 201 {object} res.Response{data=map[string]interface{}}
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	// cek apakah email sudah dipakai
	var count int64
	database.DB.Model(&model.User{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.ErrorResponse("Email already used", nil))
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	user := model.User{Name: req.Name, Email: req.Email, Password: string(hash)}
	if err := database.DB.Create(&user).Error; err != nil {
		return utils.HandleServerError(c, err)
	}

	tokens, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.Status(fiber.StatusCreated).
		JSON(res.SuccessResponse("User registered successfully", fiber.Map{
			"user": fiber.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
			"tokens": fiber.Map{
				"accessToken":  tokens.AccessToken,
				"refreshToken": tokens.RefreshToken,
			},
		}))
}

// ===================== LOGIN =====================
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body LoginRequest true "Login payload"
// @Success 200 {object} res.Response{data=map[string]interface{}}
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	var user model.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid credentials", nil))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid credentials", nil))
	}

	tokens, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("Login successful", fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"tokens": fiber.Map{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	}))
}

// ===================== REFRESH TOKEN =====================
// @Summary Refresh access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body map[string]string true "refresh token"
// @Success 200 {object} res.Response{data=map[string]string}
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/refresh [post]
func Refresh(c *fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(body); err != nil {
		return utils.HandleValidationError(c, err)
	}

	t, claims, err := utils.ParseToken(body.RefreshToken)
	if err != nil || !t.Valid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid or expired token", nil))
	}

	if typ, _ := claims["type"].(string); typ != "refresh" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid token type", nil))
	}

	subStr, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid subject", nil))
	}

	sub, err := strconv.ParseUint(subStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.ErrorResponse("Invalid subject format", nil))
	}

	tokens, err := utils.GenerateTokenPair(uint(sub))
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("Token refreshed successfully", fiber.Map{
		"accessToken":  tokens.AccessToken,
		"refreshToken": tokens.RefreshToken,
	}))
}

// ===================== LOGOUT =====================
// @Summary Logout user
// @Tags Auth
// @Produce json
// @Security Bearer
// @Success 200 {object} res.Response
// @Router /auth/logout [post]
func Logout(c *fiber.Ctx) error {
	// stateless → cukup balikan sukses
	return c.JSON(res.SuccessResponse("Logout successful", nil))
}

// ===================== CHECK TOKEN =====================
// @Summary Check token validity
// @Tags Auth
// @Produce json
// @Security Bearer
// @Success 200 {object} res.Response{data=map[string]interface{}}
// @Failure 401 {object} res.Response
// @Router /auth/check [get]
func CheckToken(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	return c.JSON(res.SuccessResponse("Token is valid", fiber.Map{
		"userID": userID,
	}))
}
