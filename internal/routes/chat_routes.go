package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"
	"QuickWork/internal/websocket"

	"github.com/gofiber/fiber/v2"
	fiberWS "github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func RegisterChatRoutes(app *fiber.App, db *gorm.DB) {
	chatRepo := repositories.NewChatRepository(db)
	chatService := services.NewChatService(chatRepo)
	chatController := controllers.NewChatController(chatService)

	// CHAT
	websocket.StartChatHub(db)

	// Middleware Upgrade WebSocket
	app.Use("/api/chat/ws", func(c *fiber.Ctx) error {
		if fiberWS.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket Handler
	app.Get("/api/chat/ws", fiberWS.New(websocket.HandleWS))

	// REST API
	app.Get("/api/chat/history", chatController.GetChatHistory)
}
