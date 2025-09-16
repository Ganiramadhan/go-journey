package controller

import (
	"go-journey/src/database"
	"go-journey/src/model"
	"go-journey/src/res"
	"go-journey/src/utils"
	"go-journey/src/validation"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// ===================== REGISTER =====================
// @Summary Register a new user
// @Description Register a new user with role 'guest' or 'admin'
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body validation.RegisterRequest true "Register payload"
// @Success 201 {object} res.Response{data=map[string]interface{}}
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var req validation.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	// TODO :: set default role to 'guest' if not provided
	if req.Role == "" {
		req.Role = "guest"
	}

	if req.Role != "guest" && req.Role != "admin" {
		return c.Status(fiber.StatusBadRequest).JSON(res.ErrorResponse("Invalid role provided", nil))
	}

	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(res.ErrorResponse("Username already used", nil))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	user := model.User{
		Username: req.Username,
		FullName: req.FullName,
		Password: string(hash),
		Role:     req.Role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return utils.HandleServerError(c, err)
	}

	tokens, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(res.SuccessResponse("User registered successfully", fiber.Map{
		"user": fiber.Map{
			"id":              user.ID,
			"username":        user.Username,
			"full_name":       user.FullName,
			"role":            user.Role,
			"register_date":   user.RegisterDate,
			"esign_id":        user.EsignID,
			"esign_status_id": user.EsignStatusID,
		},
		"tokens": fiber.Map{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	}))
}

// ===================== LOGIN =====================
// @Summary Login user
// @Description Login with username and password, returns access & refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body validation.LoginRequest true "Login payload"
// @Success 200 {object} res.Response{data=map[string]interface{}}
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var req validation.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(req); err != nil {
		return utils.HandleValidationError(c, err)
	}

	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse("Invalid credentials", nil))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse("Invalid credentials", nil))
	}

	tokens, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return utils.HandleServerError(c, err)
	}

	return c.JSON(res.SuccessResponse("Login successful", fiber.Map{
		"user": fiber.Map{
			"id":              user.ID,
			"username":        user.Username,
			"full_name":       user.FullName,
			"role":            user.Role,
			"register_date":   user.RegisterDate,
			"esign_id":        user.EsignID,
			"esign_status_id": user.EsignStatusID,
		},
		"tokens": fiber.Map{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	}))
}

// ===================== REFRESH TOKEN =====================
// @Summary Refresh access token
// @Description Use refresh token to get a new access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body validation.RefreshRequest true "Refresh token payload"
// @Success 200 {object} res.Response{data=map[string]string}
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/refresh [post]
func Refresh(c *fiber.Ctx) error {
	var body validation.RefreshRequest
	if err := c.BodyParser(&body); err != nil {
		return utils.HandleValidationError(c, err)
	}
	if err := validation.ValidateStruct(body); err != nil {
		return utils.HandleValidationError(c, err)
	}

	t, claims, err := utils.ParseToken(body.RefreshToken)
	if err != nil || !t.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse("Invalid or expired token", nil))
	}

	if typ, _ := claims["type"].(string); typ != "refresh" {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse("Invalid token type", nil))
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(res.ErrorResponse("Invalid subject", nil))
	}

	tokens, err := utils.GenerateTokenPair(sub)
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
// @Description Logout user (stateless, no server-side session)
// @Tags Auth
// @Produce json
// @Security Bearer
// @Success 200 {object} res.Response
// @Router /auth/logout [post]
func Logout(c *fiber.Ctx) error {
	return c.JSON(res.SuccessResponse("Logout successful", fiber.Map{}))
}
