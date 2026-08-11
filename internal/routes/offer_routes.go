package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterOfferRoutes(app *fiber.App, db *gorm.DB) {
	appRepo := repositories.NewApplicationRepository(db)
	appService := services.NewApplicationService(appRepo)
	appController := controllers.NewApplicationController(appService)

	app.Post("/api/applications/respond-offer",
		middleware.Protected(),
		middleware.RequireRole("student"),
		appController.RespondToOffer,
	)
}
