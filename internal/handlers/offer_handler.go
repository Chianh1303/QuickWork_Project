package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RespondOfferInput struct {
	ApplicationID uint   `json:"application_id"`
	Response      string `json:"response"` // "accept" hoặc "decline"
}

func RespondToOffer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Xác thực sinh viên đang đăng nhập
		userID := c.Locals("user_id").(float64)
		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "❌ Không tìm thấy hồ sơ sinh viên"})
		}

		var input RespondOfferInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "❌ Dữ liệu phản hồi không hợp lệ"})
		}

		// 2. Tìm đơn ứng tuyển đảm bảo đúng của sinh viên này
		var app models.Application
		if err := db.Where("id = ? AND student_id = ?", input.ApplicationID, student.ID).First(&app).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "❌ Không tìm thấy đơn ứng tuyển tương ứng"})
		}

		// Đảm bảo đơn này đang ở trạng thái đã được approved (đang có offer chờ)
		if app.Status != "approved" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "❌ Đơn ứng tuyển này hiện không có Offer nào cần xử lý"})
		}

		// 3. Cập nhật trạng thái theo quyết định của sinh viên
		if input.Response == "accept" {
			app.Status = "offer_accepted"
		} else if input.Response == "decline" {
			app.Status = "offer_declined"
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "❌ Hành động không hợp lệ"})
		}

		if err := db.Save(&app).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "❌ Lỗi hệ thống khi lưu phản hồi"})
		}

		msg := "🎉 Bạn đã đồng ý nhận offer công việc thành công!"
		if input.Response == "decline" {
			msg = "❌ Bạn đã từ chối offer công việc thành công."
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": msg,
			"data":    app,
		})
	}
}
