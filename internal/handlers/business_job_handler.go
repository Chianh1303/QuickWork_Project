package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Struct nhận dữ liệu Đăng Job từ Doanh nghiệp gửi lên
type CreateJobInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Salary      float64 `json:"salary"`
	Slots       int     `json:"slots"`
	WorkingDate string  `json:"working_date"`
}

func CreateJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Không tìm thấy hồ sơ doanh nghiệp của tài khoản này",
			})
		}
		var user models.User

		if err := db.First(&user, uint(userID)).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Không tìm thấy tài khoản",
			})
		}

		if user.Status != "approved" && user.Status != "active" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Tài khoản doanh nghiệp của bạn đang chờ duyệt hoặc đã bị từ chối.",
			})
		}
		var input CreateJobInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		if input.Title == "" || input.Description == "" || input.Location == "" || input.Salary <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Vui lòng điền đầy đủ tiêu đề, mô tả, địa điểm và lương lớn hơn 0"})
		}

		newJob := models.Job{
			BusinessID:  business.ID,
			Title:       input.Title,
			Description: input.Description,
			Location:    input.Location,
			Salary:      input.Salary,
			Slots:       input.Slots,
			Status:      "pending",
			WorkingDate: input.WorkingDate,
		}

		if err := db.Create(&newJob).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể đăng tin tuyển dụng lúc này"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Đăng tin thành công! Vui lòng chờ Admin phê duyệt", "data": newJob})
	}
}
