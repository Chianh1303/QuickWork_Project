package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterSavedJobRoutes(app *fiber.App, db *gorm.DB) {
	savedJobRepo := repositories.NewSavedJobRepository(db)
	savedJobService := services.NewSavedJobService(savedJobRepo)
	savedJobController := controllers.NewSavedJobController(savedJobService)

	saved := app.Group("/api/saved-jobs", middleware.Protected(), middleware.RequireRole("student"))
	saved.Post("/:jobId", savedJobController.ToggleSaveJob)
	saved.Get("/", savedJobController.GetSavedJobs)
}
