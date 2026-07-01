package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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

func GetEmployerApplications(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)
		var business models.Business
		if err := db.Where("user_id = ?", userID).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "❌ Không tìm thấy thông tin doanh nghiệp"})
		}

		var apps []models.Application
		err := db.
			Preload("Student").
			Preload("Student.User").
			Preload("Job").
			Preload("Job.Business").
			Preload("Job.Business.User").
			Joins("JOIN jobs ON jobs.id = applications.job_id").
			Where("jobs.business_id = ?", business.ID).
			Find(&apps).Error

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "❌ Lỗi hệ thống khi quét danh sách ứng viên",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    apps,
		})
	}
}
