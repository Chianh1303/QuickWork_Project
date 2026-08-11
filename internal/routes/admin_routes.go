package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(app *fiber.App, db *gorm.DB) {
	adminRepo := repositories.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

	admin := app.Group("/api/admin", middleware.Protected(), middleware.RequireRole("admin"))

	admin.Get("/dashboard/stats", adminController.GetAdminDashboardStats)
	admin.Get("/businesses/pending", adminController.GetPendingBusinesses)
	admin.Get("/businesses/:id", adminController.GetBusinessKYBDetail)
	admin.Put("/businesses/:id/review", adminController.ReviewBusinessKYB)
}
