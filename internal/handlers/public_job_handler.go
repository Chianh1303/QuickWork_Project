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
		if err := query.
			Preload("Business").
			Where("status = ?", "approved").
			Order("id DESC").
			Find(&jobs).Error; err != nil {
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
