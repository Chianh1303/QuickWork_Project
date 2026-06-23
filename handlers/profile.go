package handlers

import (
	"QuickWork/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 1. Struct nhận dữ liệu cập nhật cho Sinh viên
type UpdateStudentInput struct {
	FullName  string `json:"full_name"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	AvatarUrl string `json:"avatar_url"`
	Skills    string `json:"skills"` // Chuỗi JSON dạng ["Go", "React"]
	CvUrl     string `json:"cv_url"`
}

// 2. Struct nhận dữ liệu cập nhật cho Doanh nghiệp
type UpdateBusinessInput struct {
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	LogoUrl     string `json:"logo_url"`
}

// Handler cập nhật hồ sơ Sinh viên
func UpdateStudentProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Lấy user_id do Middleware Protected() bốc từ Token ra cắm vào trước đó
		userID := c.Locals("user_id").(float64) // Fiber ép kiểu số trong JWT mặc định là float64

		var input UpdateStudentInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		// Tìm hồ sơ Student dựa theo user_id
		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy hồ sơ sinh viên"})
		}

		// Cập nhật các trường dữ liệu mới vào DB
		db.Model(&student).Updates(models.Student{
			FullName:  input.FullName,
			Phone:     input.Phone,
			Gender:    input.Gender,
			AvatarUrl: input.AvatarUrl,
			Skills:    input.Skills,
			CvUrl:     input.CvUrl,
		})

		return c.JSON(fiber.Map{
			"message": "🎉 Cập nhật hồ sơ sinh viên thành công!",
			"data":    student,
		})
	}
}

// Handler cập nhật hồ sơ Doanh nghiệp
func UpdateBusinessProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var input UpdateBusinessInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Không tìm thấy hồ sơ doanh nghiệp"})
		}

		// Cập nhật thông tin doanh nghiệp
		db.Model(&business).Updates(models.Business{
			CompanyName: input.CompanyName,
			Phone:       input.Phone,
			Address:     input.Address,
			LogoUrl:     input.LogoUrl,
		})

		return c.JSON(fiber.Map{
			"message": "🎉 Cập nhật hồ sơ doanh nghiệp thành công!",
			"data":    business,
		})
	}
}