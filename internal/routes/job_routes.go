package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterJobRoutes(app *fiber.App, db *gorm.DB) {
	// Jobs Routes
	app.Post("/api/jobs", middleware.Protected(), middleware.RequireRole("business"), handlers.CreateJob(db))
	app.Get("/api/jobs", handlers.GetAvailableJobs(db)) // Công khai

	app.Put("/api/jobs/review-application", middleware.Protected(), middleware.RequireRole("business"), handlers.ReviewApplication(db))
}
