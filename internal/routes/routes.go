package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	RegisterRoutes(app, db)
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	RegisterAuthRoutes(app, db)
	RegisterProfileRoutes(app, db)
	RegisterJobRoutes(app, db)
	RegisterApplicationRoutes(app, db)
	RegisterOfferRoutes(app, db)
	RegisterAttendanceRoutes(app, db)
	RegisterWalletRoutes(app, db)
	RegisterChatRoutes(app, db)
	RegisterReviewRoutes(app, db)
	RegisterAdminRoutes(app, db)
	RegisterAIRoutes(app, db)
	RegisterNotificationRoutes(app, db)
	RegisterTicketRoutes(app, db)
}
