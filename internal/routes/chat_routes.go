package routes

import (
	"QuickWork/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func RegisterChatRoutes(app *fiber.App, db *gorm.DB) {
	// CHAT
	handlers.StartChatHub(db)

	// Middleware Upgrade WebSocket
	app.Use("/api/chat/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// REST API
	app.Get("/api/chat/history", handlers.GetChatHistory(db))

	// WebSocket
	app.Get("/api/chat/ws", websocket.New(handlers.HandleWS))
}
