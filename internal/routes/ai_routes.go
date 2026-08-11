package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAIRoutes(app *fiber.App, db *gorm.DB) {
	aiHandler := handlers.NewAIHandler(db)

	api := app.Group("/api/ai")

	// Đánh giá CV (Đã đăng nhập)
	api.Post("/evaluate-cv", middleware.Protected(), aiHandler.EvaluateCV)

	// Đánh giá độ phù hợp với Job
	api.Post("/match-job", middleware.Protected(), aiHandler.MatchJob)

	// AI Soạn bài đăng tuyển dụng (Dành cho Doanh nghiệp)
	api.Post("/generate-job", middleware.Protected(), middleware.RequireRole("business"), aiHandler.GenerateJobDescription)
}
