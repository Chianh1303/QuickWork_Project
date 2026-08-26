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
		var lockedErr *services.AccountLockedError
		if errors.As(err, &lockedErr) {
			return c.Status(fiber.StatusLocked).JSON(fiber.Map{
				"message":           lockedErr.Message,
				"remaining_seconds": lockedErr.RemainingSeconds,
			})
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

// ForgotPassword POST /api/auth/forgot-password
func (ctrl *AuthController) ForgotPassword(c *fiber.Ctx) error {
	var req dto.ForgotPasswordRequest
	if err := c.BodyParser(&req); err != nil || req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Vui lòng nhập Email hợp lệ"})
	}

	err := ctrl.authService.SendPasswordResetOTP(req.Email)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "📧 Mã xác thực OTP (6 chữ số) đã được gửi đến Email của bạn! Vui lòng kiểm tra hộp thư (Mã có hiệu lực trong 5 phút).",
	})
}

// ResetPassword POST /api/auth/reset-password
func (ctrl *AuthController) ResetPassword(c *fiber.Ctx) error {
	var req dto.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	err := ctrl.authService.VerifyPasswordResetOTP(req.Email, req.OTPCode, req.NewPassword)
	if err != nil {
		if errors.Is(err, services.ErrInvalidOTP) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "🎉 Đặt lại mật khẩu thành công! Bạn có thể đăng nhập bằng Mật khẩu mới ngay bây giờ.",
	})
}

// GoogleLogin POST /api/auth/google-login
func (ctrl *AuthController) GoogleLogin(c *fiber.Ctx) error {
	var req dto.GoogleLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu Google không hợp lệ"})
	}

	res, err := ctrl.authService.GoogleLogin(req)
	if err != nil {
		var forbiddenErr *services.AccountForbiddenError
		if errors.As(err, &forbiddenErr) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": forbiddenErr.Message})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
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
