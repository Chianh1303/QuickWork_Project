package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 🌟 ĐÃ FIX: Struct nhận dữ liệu ứng tuyển từ Sinh viên nội bộ (Thêm CoverNote)
type ApplyJobInput struct {
	JobID     uint   `json:"job_id"`
	CoverNote string `json:"cover_note"` // Nhận lời nhắn từ Nuxt 4 gửi sang
}

func GetStudentApplications(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userID := c.Locals("user_id").(float64)
		var student models.Student
		if err := db.Where("user_id = ?", userID).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "❌ Không tìm thấy hồ sơ sinh viên của tài khoản này"})
		}
		var apps []models.Application
		if err := db.
			Preload("Job").
			Preload("Job.Business").
			Preload("Job.Business.User").
			Where("student_id = ?", student.ID).
			Find(&apps).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "❌ Lỗi hệ thống khi lấy lịch sử ứng tuyển"})
		}
		return c.JSON(fiber.Map{
			"success": true,
			"data":    apps,
		})

	}
}

func ApplyJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy hồ sơ sinh viên"})
		}

		var input ApplyJobInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		var job models.Job
		if err := db.First(&job, input.JobID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Công việc này không tồn tại"})
		}

		// Kiểm tra xem đã ứng tuyển trùng chưa
		var existApp models.Application
		err := db.Where("job_id = ? AND student_id = ?", input.JobID, student.ID).First(&existApp).Error
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "⚠️ Bạn đã ứng tuyển công việc này rồi!"})
		}

		// Tiến hành lưu đơn ứng tuyển kèm Cover Note
		newApplication := models.Application{
			JobID:     input.JobID,
			StudentID: student.ID,
			Status:    "pending",
			CoverNote: input.CoverNote, // Hết lỗi undefined nhé Chanh!
		}

		if err := db.Create(&newApplication).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể nộp đơn ứng tuyển lúc này"})
		}

		return c.JSON(fiber.Map{
			"message": "🚀 Nộp đơn ứng tuyển thành công! Đang chờ Doanh nghiệp phản hồi.",
			"data":    newApplication,
		})
	}
}

func CancelApplication(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Lấy ID đơn ứng tuyển từ URL Param
		appID, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID đơn ứng tuyển không hợp lệ"})
		}

		// 2. Xác thực quyền sở hữu của sinh viên
		userID := c.Locals("user_id").(float64)
		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy hồ sơ sinh viên"})
		}

		var app models.Application
		if err := db.Where("id = ? AND student_id = ?", appID, student.ID).First(&app).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy đơn ứng tuyển này của bạn"})
		}

		// 3. Chỉ cho phép hủy khi còn ở trạng thái pending
		if app.Status != "pending" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "⚠️ Không thể hủy đơn ứng tuyển đã được doanh nghiệp xử lý!"})
		}

		// 4. Xóa đơn khỏi DB
		if err := db.Delete(&app).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể hủy đơn lúc này, vui lòng thử lại"})
		}

		return c.JSON(fiber.Map{
			"message": "❌ Đã hủy đơn ứng tuyển thành công!",
		})
	}
}

func StudentCompleteJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var input struct {
			ApplicationID uint `json:"application_id"`
		}

		if err := c.BodyParser(&input); err != nil || input.ApplicationID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "Thiếu application_id"})
		}

		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy hồ sơ sinh viên"})
		}

		var app models.Application
		if err := db.
			Preload("Job").
			Where("id = ? AND student_id = ?", input.ApplicationID, student.ID).
			First(&app).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy đơn ứng tuyển"})
		}

		if app.Status != "offer_accepted" {
			return c.Status(400).JSON(fiber.Map{"error": "Chỉ có thể hoàn thành khi đã nhận việc"})
		}

		app.Status = "student_completed"

		if err := db.Save(&app).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Không thể cập nhật trạng thái"})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "Sinh viên đã xác nhận hoàn thành công việc. Đang chờ doanh nghiệp xác nhận.",
			"data":    app,
		})
	}
}
