package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Auth Routes
	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)

	// User Check Me
	app.Get("/api/users/me", middleware.Protected(), authController.GetMe)
}
