package routes

import (
	"QuickWork/internal/clients"
	"QuickWork/internal/config"
	"QuickWork/internal/controllers"
	"QuickWork/internal/engine"
	"QuickWork/internal/middleware"
	"QuickWork/internal/parsers"
	"QuickWork/internal/queue"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAIRoutes(app *fiber.App, db *gorm.DB) {
	aiRepo := repositories.NewAIRepository(db)
	jobRepo := repositories.NewJobRepository(db)
	geminiClient := clients.NewGeminiClient()
	pdfParser := parsers.NewPDFParser()
	matchingEngine := engine.NewMatchingEngine()

	aiService := services.NewAIService(aiRepo, jobRepo, geminiClient, pdfParser, matchingEngine)
	aiController := controllers.NewAIController(aiService)

	rmqClient := queue.NewRabbitMQClient(config.RabbitMQURL)
	queue.RegisterWorkers(rmqClient, aiService.ProcessBackgroundCVEvaluation)

	api := app.Group("/api/ai")

	// Đánh giá CV (Đã đăng nhập)
	api.Post("/evaluate-cv", middleware.Protected(), aiController.EvaluateCV)
	api.Get("/latest-cv-evaluation", middleware.Protected(), middleware.RequireRole("student"), aiController.GetLatestCVEvaluation)

	// AI Gợi ý công việc cho Sinh viên
	api.Get("/recommended-jobs", middleware.Protected(), middleware.RequireRole("student"), aiController.GetRecommendedJobs)

	// Đánh giá độ phù hợp với Job
	api.Post("/match-job", middleware.Protected(), aiController.MatchJob)

	// AI Soạn bài đăng tuyển dụng (Dành cho Doanh nghiệp)
	api.Post("/generate-job", middleware.Protected(), middleware.RequireRole("business"), aiController.GenerateJobDescription)
}
