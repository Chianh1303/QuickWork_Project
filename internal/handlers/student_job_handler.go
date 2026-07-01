package handlers

import (
	"QuickWork/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAvailableJobs(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Đón đầu toàn bộ 5 tham số từ URL Query String
		search := c.Query("search")
		location := c.Query("location")
		category := c.Query("category")
		jobType := c.Query("job_type")
		maxSalary := c.Query("max_salary")

		var jobs []models.Job
		query := db.Model(&models.Job{})

		// 2. Lọc theo Từ khóa (Title)
		if search != "" {
			query = query.Where("LOWER(title) LIKE LOWER(?)", "%"+search+"%")
		}

		// 3. Lọc theo Địa điểm
		if location != "" && location != "all" {
			query = query.Where("LOWER(location) LIKE LOWER(?)", "%"+location+"%")
		}

		// 4. Lọc theo Ngành nghề (Category)
		if category != "" && category != "all" {
			query = query.Where("LOWER(category) = LOWER(?)", category)
		}

		// 5. Lọc theo Hình thức (Job Type)
		if jobType != "" && jobType != "all" {
			query = query.Where("LOWER(job_type) = LOWER(?)", jobType)
		}

		// 6. Lọc theo Mức lương tối thiểu
		if maxSalary != "" {
			query = query.Where("salary >= ?", maxSalary)
		}

		// Thực thi truy vấn đưa ra danh sách
		// 🌟 Chèn thêm .Where vào giữa để lọc đúng các Job đã được Admin duyệt
		if err := query.Where("status = ?", "approved").Order("id DESC").Find(&jobs).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Không thể lấy danh sách công việc toàn diện",
			})
		}

		return c.JSON(fiber.Map{
			"message": "🎨 Lấy danh sách công việc thành công!",
			"data":    jobs,
		})
	}
}

// 🌟 ĐÃ FIX: Struct nhận dữ liệu ứng tuyển từ Sinh viên nội bộ (Thêm CoverNote)
type ApplyJobInput struct {
	JobID     uint   `json:"job_id"`
	CoverNote string `json:"cover_note"` // Nhận lời nhắn từ Nuxt 4 gửi sang
}

type RespondOfferInput struct {
	ApplicationID uint   `json:"application_id"`
	Response      string `json:"response"` // "accept" hoặc "decline"
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
