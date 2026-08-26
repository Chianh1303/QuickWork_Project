package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterProfileRoutes(app *fiber.App, db *gorm.DB) {
	profileRepo := repositories.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileController := controllers.NewProfileController(profileService)

	// Profile Routes
	// --- Đối với Sinh viên ---
	app.Get("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), profileController.GetStudentProfile)
	app.Put("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), profileController.UpdateStudentProfile)

	// --- Đối với Doanh nghiệp ---
	app.Get("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), profileController.GetBusinessProfile)
	app.Put("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), profileController.UpdateBusinessProfile)
}
