package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterOfferRoutes(app *fiber.App, db *gorm.DB) {
	// Đặt cùng nhóm với các route của Student nhé Chanh
	app.Post("/api/applications/respond-offer",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.RespondToOffer(db),
	)
}
