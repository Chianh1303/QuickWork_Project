package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 🌟 ĐÃ NÂNG CẤP STRUCT: Nhận thêm các trường Offer từ Nuxt gửi lên
type ReviewApplicationInput struct {
	ApplicationID  uint   `json:"application_id"`
	Status         string `json:"status"` // "approved" hoặc "rejected"
	OfferSalary    string `json:"offer_salary"`
	OfferStartDate string `json:"offer_start_date"`
	OfferMessage   string `json:"offer_message"`
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

// 🌟 ĐÃ NÂNG CẤP HÀM: Vừa bảo mật vừa hỗ trợ gửi Offer đi kèm (B6)
func ReviewApplication(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Lấy user_id từ Token để xác định Doanh nghiệp đang đăng nhập
		userID := c.Locals("user_id").(float64)

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy tài khoản Doanh nghiệp"})
		}

		var input ReviewApplicationInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		if input.Status != "approved" && input.Status != "rejected" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Trạng thái không hợp lệ"})
		}

		var application models.Application
		if err := db.First(&application, input.ApplicationID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy đơn ứng tuyển này"})
		}

		var job models.Job
		if err := db.First(&job, application.JobID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy tin tuyển dụng này"})
		}

		// Kiểm tra quyền sở hữu bài đăng (Giữ nguyên logic bảo mật cực tốt của Chanh)
		if job.BusinessID != business.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Bạn không có quyền chỉnh sửa đơn ứng tuyển này"})
		}

		// Cập nhật trạng thái cốt lõi
		application.Status = input.Status

		// ✉️ Nếu duyệt chấp nhận, lưu kèm trọn bộ thông tin Offer mới vào DB
		if input.Status == "approved" {
			application.OfferSalary = input.OfferSalary
			application.OfferStartDate = input.OfferStartDate
			application.OfferMessage = input.OfferMessage
		}

		if err := db.Save(&application).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Lỗi hệ thống khi cập nhật phản hồi"})
		}

		msg := "🎉 Chấp nhận hồ sơ ứng viên và gửi kèm thông tin Offer thành công!"
		if input.Status == "rejected" {
			msg = "❌ Đã từ chối đơn ứng tuyển thành công."
		}

		return c.JSON(fiber.Map{
			"message": msg,
			"data":    application,
		})
	}
}

func BusinessCompleteJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var input struct {
			ApplicationID uint `json:"application_id"`
		}

		if err := c.BodyParser(&input); err != nil || input.ApplicationID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "Thiếu application_id"})
		}

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy hồ sơ doanh nghiệp"})
		}

		var app models.Application
		if err := db.
			Preload("Job").
			Preload("Student").
			Where("id = ?", input.ApplicationID).
			First(&app).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy đơn ứng tuyển"})
		}

		if app.Job.BusinessID != business.ID {
			return c.Status(403).JSON(fiber.Map{"error": "Bạn không có quyền xác nhận đơn này"})
		}

		if app.Status != "student_completed" {
			return c.Status(400).JSON(fiber.Map{"error": "Sinh viên chưa xác nhận hoàn thành"})
		}

		err := db.Transaction(func(tx *gorm.DB) error {
			app.Status = "paid"

			if err := tx.Save(&app).Error; err != nil {
				return err
			}

			var studentUser models.User
			if err := tx.First(&studentUser, app.Student.UserID).Error; err != nil {
				return err
			}

			var wallet models.Wallet
			if err := tx.Where("user_id = ?", studentUser.ID).First(&wallet).Error; err != nil {
				wallet = models.Wallet{
					UserID:  studentUser.ID,
					Balance: 0,
				}
				if err := tx.Create(&wallet).Error; err != nil {
					return err
				}
			}

			amount := app.Job.Salary
			wallet.Balance += amount

			if err := tx.Save(&wallet).Error; err != nil {
				return err
			}

			transaction := models.WalletTransaction{
				WalletID:      wallet.ID,
				Type:          "salary",
				Amount:        amount,
				Description:   "Lương công việc: " + app.Job.Title,
				ReferenceID:   app.ID,
				ReferenceType: "application",
			}

			return tx.Create(&transaction).Error
		})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Không thể xác nhận và giải ngân"})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "Doanh nghiệp đã xác nhận hoàn thành và giải ngân lương.",
			"data":    app,
		})
	}
}
