package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterNotificationRoutes(app *fiber.App, db *gorm.DB) {
	repo := repositories.NewNotificationRepository(db)
	service := services.NewNotificationService(repo)
	ctrl := controllers.NewNotificationController(service)

	// Protected routes
	api := app.Group("/api/notifications", middleware.Protected())
	api.Get("/", ctrl.GetUserNotifications)
	api.Patch("/read-all", ctrl.MarkAllAsRead)
	api.Post("/read-all", ctrl.MarkAllAsRead)
	api.Patch("/:id/read", ctrl.MarkAsRead)
	api.Post("/:id/read", ctrl.MarkAsRead)
}
