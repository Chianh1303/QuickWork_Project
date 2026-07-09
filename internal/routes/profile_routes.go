package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterProfileRoutes(app *fiber.App, db *gorm.DB) {
	// Profile Routes
	// --- Đối với Sinh viên ---
	app.Get("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), handlers.GetStudentProfile(db))
	app.Put("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), handlers.UpdateStudentProfile(db))

	// --- Đối với Doanh nghiệp ---
	app.Get("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), handlers.GetBusinessProfile(db))
	app.Put("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), handlers.UpdateBusinessProfile(db))
}
