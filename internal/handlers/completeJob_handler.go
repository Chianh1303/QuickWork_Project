package handlers

import (
	"QuickWork/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func StudentCompleteJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token không hợp lệ"})
		}

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
		if err := db.Where("id = ? AND student_id = ?", input.ApplicationID, student.ID).First(&app).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy đơn ứng tuyển"})
		}

		if app.Status != "offer_accepted" && app.Status != "working" {
			return c.Status(400).JSON(fiber.Map{"error": "Công việc chưa ở trạng thái có thể xác nhận hoàn thành"})
		}

		app.StudentCompleted = true
		app.Status = "student_completed"

		if err := db.Save(&app).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Không thể cập nhật trạng thái"})
		}

		return c.JSON(fiber.Map{
			"message": "Bạn đã xác nhận hoàn thành công việc. Đang chờ doanh nghiệp xác nhận.",
			"data":    app,
		})
	}
}

func BusinessCompleteJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token không hợp lệ"})
		}

		var input struct {
			ApplicationID uint `json:"application_id"`
		}

		if err := c.BodyParser(&input); err != nil || input.ApplicationID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "Thiếu application_id"})
		}

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy doanh nghiệp"})
		}

		var app models.Application
		if err := db.
			Preload("Job").
			Where("applications.id = ?", input.ApplicationID).
			First(&app).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy đơn ứng tuyển"})
		}

		if app.Job.BusinessID != business.ID {
			return c.Status(403).JSON(fiber.Map{"error": "Bạn không có quyền xác nhận đơn này"})
		}

		if !app.StudentCompleted {
			return c.Status(400).JSON(fiber.Map{"error": "Sinh viên chưa xác nhận hoàn thành công việc"})
		}

		if app.BusinessCompleted || app.Status == "paid" {
			return c.JSON(fiber.Map{
				"message": "Công việc này đã được doanh nghiệp xác nhận và giải ngân trước đó.",
				"data":    app,
			})
		}

		now := time.Now()

app.BusinessCompleted = true
app.CompletedAt = &now
app.PaidAt = &now
app.PaymentStatus = "paid"
app.Status = "paid"

var student models.Student
if err := db.First(&student, app.StudentID).Error; err != nil {
	return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy sinh viên để giải ngân"})
}

salaryAmount := app.Job.Salary
if app.OfferSalary != "" {
	parsedSalary, err := strconv.ParseFloat(app.OfferSalary, 64)
	if err == nil && parsedSalary > 0 {
		salaryAmount = parsedSalary
	}
}

err := db.Transaction(func(tx *gorm.DB) error {
	if err := tx.Save(&app).Error; err != nil {
		return err
	}

	var studentWallet models.Wallet
	if err := tx.Where("user_id = ?", student.UserID).First(&studentWallet).Error; err != nil {
		studentWallet = models.Wallet{
			UserID:  student.UserID,
			Balance: 0,
		}
		if err := tx.Create(&studentWallet).Error; err != nil {
			return err
		}
	}

	studentWallet.Balance += salaryAmount

	if err := tx.Save(&studentWallet).Error; err != nil {
		return err
	}

	studentTransaction := models.WalletTransaction{
		WalletID:      studentWallet.ID,
		Type:          "salary",
		Amount:        salaryAmount,
		Description:   "Nhận lương từ công việc: " + app.Job.Title,
		ReferenceID:   app.ID,
		ReferenceType: "application",
	}

	if err := tx.Create(&studentTransaction).Error; err != nil {
		return err
	}

	return nil
})

if err != nil {
	return c.Status(500).JSON(fiber.Map{
		"error": "Không thể giải ngân lương",
	})
}

		return c.JSON(fiber.Map{
			"message": "Doanh nghiệp đã xác nhận hoàn thành. Hệ thống đã giả lập giải ngân lương.",
			"data":    app,
		})
	}
}
