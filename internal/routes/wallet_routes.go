package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterWalletRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/wallet/me",
		middleware.Protected(),
		handlers.GetMyWallet(db),
	)
}
