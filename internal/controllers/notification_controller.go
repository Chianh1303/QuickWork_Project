package controllers

import (
	"strconv"

	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	notifService services.NotificationService
}

func NewNotificationController(notifService services.NotificationService) *NotificationController {
	return &NotificationController{notifService: notifService}
}

// GetUserNotifications GET /api/notifications
func (ctrl *NotificationController) GetUserNotifications(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	notifs, unreadCount, err := ctrl.notifService.GetUserNotifications(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Không thể lấy danh sách thông báo"})
	}

	return c.JSON(fiber.Map{
		"data":         notifs,
		"unread_count": unreadCount,
	})
}

// MarkAsRead PATCH /api/notifications/:id/read
func (ctrl *NotificationController) MarkAsRead(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	notifID, err := c.ParamsInt("id")
	if err != nil || notifID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID thông báo không hợp lệ"})
	}

	if err := ctrl.notifService.MarkAsRead(uint(notifID), userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Không thể cập nhật thông báo"})
	}

	return c.JSON(fiber.Map{"message": "Đã đánh dấu đã đọc"})
}

// MarkAllAsRead PATCH /api/notifications/read-all
func (ctrl *NotificationController) MarkAllAsRead(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	if err := ctrl.notifService.MarkAllAsRead(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Không thể đánh dấu tất cả"})
	}

	return c.JSON(fiber.Map{"message": "Đã đánh dấu tất cả thông báo là đã đọc"})
}
