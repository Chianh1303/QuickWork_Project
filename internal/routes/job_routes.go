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

	notifRepo := repositories.NewNotificationRepository(db)
	notifService := services.NewNotificationService(notifRepo)
	appRepo := repositories.NewApplicationRepository(db)
	appService := services.NewApplicationService(appRepo, notifService)
	appController := controllers.NewApplicationController(appService)

	// Static routes MUST come before parameterized routes (:id)
	app.Get("/api/jobs", jobController.GetAvailableJobs) // Công khai cho sinh viên
	app.Get("/api/jobs/business/my-jobs", middleware.Protected(), middleware.RequireRole("business"), jobController.GetMyBusinessJobs)
	app.Put("/api/jobs/review-application", middleware.Protected(), middleware.RequireRole("business"), appController.ReviewApplication)

	// Business Parameterized Routes (:id)
	app.Post("/api/jobs", middleware.Protected(), middleware.RequireRole("business"), jobController.CreateJob)
	app.Put("/api/jobs/:id", middleware.Protected(), middleware.RequireRole("business"), jobController.UpdateJob)
	app.Patch("/api/jobs/:id/status", middleware.Protected(), middleware.RequireRole("business"), jobController.ToggleJobStatus)
	app.Delete("/api/jobs/:id", middleware.Protected(), middleware.RequireRole("business"), jobController.DeleteJob)
}
