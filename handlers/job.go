package handlers

import (
	"QuickWork/models"

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
		// 1. Lấy user_id (ID tài khoản gốc) từ mã Token do Middleware cấp
		userID := c.Locals("user_id").(float64)

		// 2. Tìm ID Hồ sơ Doanh nghiệp (business_id) tương ứng với user_id này
		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Không tìm thấy hồ sơ doanh nghiệp của tài khoản này",
			})
		}

		// 3. Parse dữ liệu body gửi lên
		var input CreateJobInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		// Validation cơ bản
		if input.Title == "" || input.Description == "" || input.Location == "" || input.Salary <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Vui lòng điền đầy đủ tiêu đề, mô tả, địa điểm và lương lớn hơn 0"})
		}

		// 4. Khởi tạo đối tượng Job mới theo cấu trúc bảng chuẩn của bạn
		newJob := models.Job{
			BusinessID:  business.ID, // Gán chuẩn khóa ngoại nối sang hồ sơ Doanh nghiệp
			Title:       input.Title,
			Description: input.Description,
			Location:    input.Location,
			Salary:      input.Salary,
			Slots:       input.Slots,
			Status:      "pending", // Chờ duyệt theo đúng DBML thiết kế
			WorkingDate: input.WorkingDate,
		}

		// 5. Lưu xuống Database
		if err := db.Create(&newJob).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể đăng tin tuyển dụng lúc này"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Đăng tin thành công! Vui lòng chờ Admin phê duyệt", "data":newJob})
	}
}