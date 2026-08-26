package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterReviewRoutes(app *fiber.App, db *gorm.DB) {
	reviewRepo := repositories.NewReviewRepository(db)
	reviewService := services.NewReviewService(reviewRepo)
	reviewController := controllers.NewReviewController(reviewService)

	app.Get(
		"/api/reviews/user/:userId",
		middleware.Protected(),
		reviewController.GetReviewsByUser,
	)
	app.Get(
		"/api/reviews/application/:applicationId",
		middleware.Protected(),
		reviewController.GetReviewsByApplication,
	)
	app.Post(
		"/api/reviews",
		middleware.Protected(),
		reviewController.CreateReview,
	)
}
