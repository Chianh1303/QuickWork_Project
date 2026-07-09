package handlers

import (
	"QuickWork/internal/models"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// [GET] API Lấy profile hiện tại của Doanh nghiệp
func GetBusinessProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Không tìm thấy hồ sơ doanh nghiệp"})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    business,
		})
	}
}

// [PUT] API Cập nhật profile Doanh nghiệp + Upload Logo thực tế
func UpdateBusinessProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var business models.Business
		if err := db.Where("user_id = ?", uint(userID)).First(&business).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Không tìm thấy hồ sơ doanh nghiệp"})
		}

		// Đọc dữ liệu text từ Form-data
		business.CompanyName = c.FormValue("company_name", business.CompanyName)
		business.Phone = c.FormValue("phone", business.Phone)
		business.Address = c.FormValue("address", business.Address)

		// Xử lý tải file thực tế: Logo Công ty
		logoFile, err := c.FormFile("logo")
		if err == nil {
			logoName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(logoFile.Filename))
			logoPath := filepath.Join("./uploads/avatars", logoName) // Dùng chung thư mục lưu ảnh đại diện

			if err := c.SaveFile(logoFile, logoPath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Lỗi khi lưu file Logo công ty"})
			}
			business.LogoUrl = fmt.Sprintf("http://localhost:3000/uploads/avatars/%s", logoName)
		}

		if err := db.Save(&business).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Không thể cập nhật DB doanh nghiệp"})
		}

		return c.JSON(fiber.Map{
			"message": "🎉 Cập nhật hồ sơ doanh nghiệp thành công!",
			"data":    business,
		})
	}
}
