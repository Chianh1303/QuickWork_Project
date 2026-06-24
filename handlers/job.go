package handlers

import (
	"QuickWork/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReviewApplicationInput struct {
	ApplicationID uint   `json:"application_id"`
	Status        string `json:"status"`     // "accepted" hoặc "rejected"
}

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

func ReviewApplication(db *gorm.DB) fiber.Handler {
	return func (c *fiber.Ctx) error {
		// 1. Lấy user_id từ Token để xác định Doanh nghiệp đang đăng nhập
		userID := c.Locals("user_id").(float64)

		var business models.Business
		if err := db.Where("user_id = ?",uint(userID)).First(&business).Error; err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Không tìm thấy tài khoản Doanh nghiệp"})
		}
		var input ReviewApplicationInput
		if err :=c.BodyParser(&input); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Dữ liệu không hợp lệ"})
		}

		if input.Status != "approved" && input.Status != "rejected"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Trạng thái không hợp lệ"})
		}
		var application models.Application
		if err :=db.First(&application,input.ApplicationID).Error; err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Không tìm thấy đơn ứng tuyển này"})
		}
		var job models.Job
		if err := db.First(&job,application.JobID).Error; err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Không tìm thấy tin tuyển dụng này"})
		}
		if job.BusinessID != business.ID{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message":"Bạn không có quyền chỉnh sửa đơn ứng tuyển này"})
		}
		application.Status=input.Status
		if err := db.Save(&application).Error; err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Lỗi hệ thống"})
		}
		return c.JSON(fiber.Map{
			"message": "🎉 Xử lý đơn ứng tuyển thành công!",
			"data" : application,
		})
	}
}