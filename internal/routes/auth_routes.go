package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	// Auth Routes
	app.Post("/api/auth/register", handlers.HandleRegister(db))
	app.Post("/api/auth/login", handlers.HandleLogin(db))

	// User Check Me
	app.Get("/api/users/me", middleware.Protected(), func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")

		return c.JSON(fiber.Map{
			"message": "🔓 Bạn đã vượt qua trạm kiểm soát bảo mật thành công!",
			"user_id": userID,
			"role":    role,
		})
	})
}
