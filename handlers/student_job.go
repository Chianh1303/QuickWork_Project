package handlers

import (
	"QuickWork/models"
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
        maxSalary := c.Query("max_salary") // Lọc các công việc có lương từ mức này trở lên chẳng hạn

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
// 4. Lọc theo Ngành nghề (Category) - Chỉ ép lọc khi param truyền lên hợp lệ
if category != "" && category != "all" {
    query = query.Where("LOWER(category) = LOWER(?)", category)
}

// 5. Lọc theo Hình thức (Job Type) - Chỉ ép lọc khi param truyền lên hợp lệ
if jobType != "" && jobType != "all" {
    query = query.Where("LOWER(job_type) = LOWER(?)", jobType)
}

        // 6. Lọc theo Mức lương tối thiểu (Sinh viên muốn tìm việc từ X triệu trở lên)
        if maxSalary != "" {
            query = query.Where("salary >= ?", maxSalary)
        }

        // Thực thi truy vấn đưa ra danh sách
        if err := query.Order("id DESC").Find(&jobs).Error; err != nil {
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

// Struct nhận dữ liệu ứng tuyển từ Sinh viên
type ApplyJobInput struct {
	JobID uint `json:"job_id"`
}

// 2. API Sinh viên bấm Ứng tuyển (Bắt buộc Token + Quyền student)
func ApplyJob(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Lấy user_id gốc từ Token
		userID := c.Locals("user_id").(float64)

		// Tìm hồ sơ Student tương ứng để lấy StudentID
		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy hồ sơ sinh viên"})
		}

		var input ApplyJobInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		// Kiểm tra xem Job đó có tồn tại không
		var job models.Job
		if err := db.First(&job, input.JobID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Công việc này không tồn tại hoặc đã bị xóa"})
		}

		// Kiểm tra xem sinh viên này đã ứng tuyển job này trước đó chưa (Tránh ứng tuyển trùng)
		var existApp models.Application
		err := db.Where("job_id = ? AND student_id = ?", input.JobID, student.ID).First(&existApp).Error
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "⚠️ Bạn đã ứng tuyển công việc này rồi!"})
		}

		// Tiến hành lưu lượt ứng tuyển mới
		newApplication := models.Application{
			JobID:     input.JobID,
			StudentID: student.ID,
			Status:    "applied", // Trạng thái mặc định vừa nộp
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