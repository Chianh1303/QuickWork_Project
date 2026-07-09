package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterReviewRoutes(app *fiber.App, db *gorm.DB) {
	app.Get(
		"/api/reviews/user/:userId",
		middleware.Protected(),
		handlers.GetReviewsByUser(db),
	)
	app.Get(
		"/api/reviews/application/:applicationId",
		middleware.Protected(),
		handlers.GetReviewsByApplication(db),
	)
	app.Post(
		"/api/reviews",
		middleware.Protected(),
		handlers.CreateReview(db),
	)
}
