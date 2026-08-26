package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterWalletRoutes(app *fiber.App, db *gorm.DB) {
	walletRepo := repositories.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo)
	walletController := controllers.NewWalletController(walletService)

	app.Get("/api/wallet/me",
		middleware.Protected(),
		walletController.GetMyWallet,
	)
}
