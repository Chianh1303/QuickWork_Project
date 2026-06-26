package handlers

import (
    "QuickWork/models"
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

// 🌟 ĐÃ NÂNG CẤP HÀM: Vừa bảo mật vừa hỗ trợ gửi Offer đi kèm (B6)
func ReviewApplication(db *gorm.DB) fiber.Handler {
    return func (c *fiber.Ctx) error {
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