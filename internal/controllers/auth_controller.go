package controllers

import (
	"errors"
	"net/http"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register POST /api/auth/register
func (ctrl *AuthController) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	err := ctrl.authService.Register(req)
	if err != nil {
		if errors.Is(err, services.ErrEmptyFields) || errors.Is(err, services.ErrInvalidRole) || errors.Is(err, services.ErrStudentNameRequired) || errors.Is(err, services.ErrBusinessFieldsRequired) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrEmailExists) || errors.Is(err, services.ErrTaxCodeExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "🎉 Đăng ký tài khoản thành công!",
	})
}

// Login POST /api/auth/login
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	res, err := ctrl.authService.Login(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}
		var forbiddenErr *services.AccountForbiddenError
		if errors.As(err, &forbiddenErr) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": forbiddenErr.Message,
				"status":  forbiddenErr.Status,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(res)
}

// GetMe GET /api/users/me
func (ctrl *AuthController) GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	role := c.Locals("role")

	return c.JSON(fiber.Map{
		"message": "🔓 Bạn đã vượt qua trạm kiểm soát bảo mật thành công!",
		"user_id": userID,
		"role":    role,
	})
}
