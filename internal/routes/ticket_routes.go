package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterTicketRoutes(app *fiber.App, db *gorm.DB) {
	ticketRepo := repositories.NewTicketRepository(db)
	appRepo := repositories.NewApplicationRepository(db)
	notifRepo := repositories.NewNotificationRepository(db)
	notifService := services.NewNotificationService(notifRepo)
	ticketService := services.NewTicketService(ticketRepo, appRepo, notifService)
	ticketController := controllers.NewTicketController(ticketService)

	app.Post("/api/tickets", middleware.Protected(), ticketController.CreateTicket)
	app.Post("/api/tickets/upload-evidence", middleware.Protected(), ticketController.UploadEvidence)
	app.Post("/api/tickets/:id/reappeal", middleware.Protected(), ticketController.ReappealTicket)
	app.Get("/api/tickets/my-tickets", middleware.Protected(), ticketController.GetUserTickets)
}
