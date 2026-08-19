package routes

import (
	"QuickWork/internal/cache"
	"QuickWork/internal/config"
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/queue"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	authRepo := repositories.NewAuthRepository(db)
	emailSvc := services.NewEmailService()
	redisClient := cache.NewRedisCache()
	rmqClient := queue.NewRabbitMQClient(config.RabbitMQURL)

	authService := services.NewAuthService(authRepo, emailSvc, redisClient, rmqClient)
	authController := controllers.NewAuthController(authService)

	// Register RabbitMQ Workers
	queue.RegisterWorkers(rmqClient, nil, emailSvc.SendOTPEmail)

	// Auth Routes
	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)
	app.Post("/api/auth/forgot-password", authController.ForgotPassword)
	app.Post("/api/auth/reset-password", authController.ResetPassword)
	app.Post("/api/auth/google-login", authController.GoogleLogin)

	// User Check Me
	app.Get("/api/users/me", middleware.Protected(), authController.GetMe)
}
