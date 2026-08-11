package routes

import (
	"QuickWork/internal/clients"
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/parsers"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAIRoutes(app *fiber.App, db *gorm.DB) {
	aiRepo := repositories.NewAIRepository(db)
	geminiClient := clients.NewGeminiClient()
	pdfParser := parsers.NewPDFParser()
	aiService := services.NewAIService(aiRepo, geminiClient, pdfParser)
	aiController := controllers.NewAIController(aiService)

	api := app.Group("/api/ai")

	// Đánh giá CV (Đã đăng nhập)
	api.Post("/evaluate-cv", middleware.Protected(), aiController.EvaluateCV)

	// Đánh giá độ phù hợp với Job
	api.Post("/match-job", middleware.Protected(), aiController.MatchJob)

	// AI Soạn bài đăng tuyển dụng (Dành cho Doanh nghiệp)
	api.Post("/generate-job", middleware.Protected(), middleware.RequireRole("business"), aiController.GenerateJobDescription)
}
