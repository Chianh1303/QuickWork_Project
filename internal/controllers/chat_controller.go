package controllers

import (
	"strconv"

	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	chatService services.ChatService
}

func NewChatController(chatService services.ChatService) *ChatController {
	return &ChatController{chatService: chatService}
}

// GetChatHistory GET /api/chat/history
func (ctrl *ChatController) GetChatHistory(c *fiber.Ctx) error {
	appIDStr := c.Query("application_id")
	appID, err := strconv.Atoi(appIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "application_id invalid",
		})
	}

	messages, err := ctrl.chatService.GetChatHistory(uint(appID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(messages)
}
