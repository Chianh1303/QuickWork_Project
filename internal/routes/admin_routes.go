package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(app *fiber.App, db *gorm.DB) {
	admin := app.Group("/api/admin", middleware.Protected(), middleware.RequireRole("admin"))

	admin.Get("/businesses/pending", handlers.GetPendingBusinesses(db))
}
