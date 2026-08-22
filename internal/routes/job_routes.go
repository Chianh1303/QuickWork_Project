package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterJobRoutes(app *fiber.App, db *gorm.DB) {
	jobRepo := repositories.NewJobRepository(db)
	jobService := services.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)

	appRepo := repositories.NewApplicationRepository(db)
	appService := services.NewApplicationService(appRepo)
	appController := controllers.NewApplicationController(appService)

	// Jobs Routes
	app.Get("/api/jobs", jobController.GetAvailableJobs) // Công khai cho sinh viên

	// Business Protected Routes
	app.Post("/api/jobs", middleware.Protected(), middleware.RequireRole("business"), jobController.CreateJob)
	app.Get("/api/jobs/business/my-jobs", middleware.Protected(), middleware.RequireRole("business"), jobController.GetMyBusinessJobs)
	app.Put("/api/jobs/:id", middleware.Protected(), middleware.RequireRole("business"), jobController.UpdateJob)
	app.Patch("/api/jobs/:id/status", middleware.Protected(), middleware.RequireRole("business"), jobController.ToggleJobStatus)
	app.Delete("/api/jobs/:id", middleware.Protected(), middleware.RequireRole("business"), jobController.DeleteJob)

	app.Put("/api/jobs/review-application", middleware.Protected(), middleware.RequireRole("business"), appController.ReviewApplication)
}
