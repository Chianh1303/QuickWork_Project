package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
	"time"
)

// ⚡ 1. Handler CHECK-IN: POST /api/attendance/check-in
func CheckIn(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 🔒 Kiểm tra an toàn sự tồn tại của locals trước khi ép kiểu
		userIDVal := c.Locals("user_id")
		if userIDVal == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Hết phiên đăng nhập, vui lòng đăng nhập lại!"})
		}

		// Kiểm tra xem kiểu dữ liệu có đúng là uint không để tránh Panic sập 502
		userID, ok := userIDVal.(float64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dữ liệu User ID không hợp lệ"})
		}

		// Tìm kiếm hồ sơ sinh viên
		var student models.Student
		if err := db.Where("user_id = ?", userID).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tài khoản của bạn chưa cập nhật hồ sơ sinh viên!"})
		}

		var input struct {
			JobID uint `json:"job_id"`
		}
		if err := c.BodyParser(&input); err != nil || input.JobID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Thiếu mã công việc (Job ID)"})
		}

		// Kiểm tra ca làm trùng lặp chưa đóng
		var existingAttendance models.Attendance
		err := db.Where("student_id = ? AND status = ?", student.ID, "working").First(&existingAttendance).Error
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bạn đang có một ca làm việc chưa kết thúc. Vui lòng Check-out ca cũ trước!",
			})
		}

		// Kiểm tra trạng thái đơn ứng tuyển (Chấp nhận cả 'offer_accepted' lẫn 'offer accepted')
		var app models.Application
		err = db.Where("student_id = ? AND job_id = ?", student.ID, input.JobID).First(&app).Error
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Bạn chưa ứng tuyển vào công việc này!"})
		}

		// Chuẩn hóa chuỗi trạng thái từ DB ra để so sánh, loại bỏ hoàn toàn lỗi SQL LOWER()
		statusNorm := strings.ToLower(strings.ReplaceAll(app.Status, "_", " "))
		if statusNorm != "offer accepted" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Công việc này chưa được trúng tuyển hoặc chưa kích hoạt ca làm!"})
		}

		// Thực hiện ghi nhận chấm công
		now := time.Now()
		attendance := models.Attendance{
			StudentID:   student.ID,
			JobID:       input.JobID,
			CheckInTime: &now,
			Status:      "working",
		}

		if err := db.Create(&attendance).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Lỗi hệ thống, không thể lưu dữ liệu chấm công"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "⚡ Check-in đầu ca thành công! Chúc bạn làm việc vui vẻ.",
			"data":    attendance,
		})
	}
}

// 🛑 2. Handler CHECK-OUT: POST /api/attendance/check-out
func CheckOut(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDVal := c.Locals("user_id")
		if userIDVal == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Hết phiên đăng nhập!"})
		}
		userID, ok := userIDVal.(float64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dữ liệu User ID không hợp lệ"})
		}

		var student models.Student
		if err := db.Where("user_id = ?", userID).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Không tìm thấy hồ sơ sinh viên"})
		}

		var input struct {
			JobID uint `json:"job_id"`
		}
		if err := c.BodyParser(&input); err != nil || input.JobID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Thiếu mã công việc (Job ID)"})
		}

		var attendance models.Attendance
		err := db.Where("student_id = ? AND job_id = ? AND status = ?", student.ID, input.JobID, "working").First(&attendance).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Không tìm thấy ca làm việc nào đang mở cần Check-out"})
		}

		now := time.Now()
		attendance.CheckOutTime = &now
		attendance.Status = "completed"

		if err := db.Save(&attendance).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Không thể cập nhật thông tin ra ca"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "🛑 Check-out ra ca thành công! Hệ thống đã ghi nhận ca làm của bạn.",
			"data":    attendance,
		})
	}
}
