package handlers

import (
	"QuickWork/internal/models"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// [GET] API Lấy profile hiện tại của Sinh viên (Để hiển thị lên FE Nuxt)
func GetStudentProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Không tìm thấy hồ sơ sinh viên"})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    student,
		})
	}
}

// [PUT] API Cập nhật profile Sinh viên + Upload file vật lý (Avatar & CV)
func UpdateStudentProfile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(float64)

		var student models.Student
		if err := db.Where("user_id = ?", uint(userID)).First(&student).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Không tìm thấy hồ sơ sinh viên"})
		}

		// Đọc các trường chữ từ Form-data gửi lên (Nếu rỗng thì giữ nguyên giá trị cũ trong DB)
		student.FullName = c.FormValue("full_name", student.FullName)
		student.Phone = c.FormValue("phone", student.Phone)
		student.Gender = c.FormValue("gender", student.Gender)
		student.Skills = c.FormValue("skills", student.Skills) // Nhận chuỗi JSON dạng ["Go", "React"]

		// Xử lý tải file thực tế: Ảnh đại diện (Avatar)
		avatarFile, err := c.FormFile("avatar")
		if err == nil {
			avatarName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(avatarFile.Filename))
			avatarPath := filepath.Join("./uploads/avatars", avatarName)

			if err := c.SaveFile(avatarFile, avatarPath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Lỗi khi lưu file ảnh đại diện"})
			}
			student.AvatarUrl = fmt.Sprintf("http://localhost:3000/uploads/avatars/%s", avatarName)
		}

		// Xử lý tải file thực tế: Hồ sơ (CV PDF)
		cvFile, err := c.FormFile("cv")
		if err == nil {
			if filepath.Ext(cvFile.Filename) != ".pdf" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Hồ sơ đính kèm bắt buộc phải là định dạng file PDF"})
			}

			cvName := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(cvFile.Filename))
			cvPath := filepath.Join("./uploads/cvs", cvName)

			if err := c.SaveFile(cvFile, cvPath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Lỗi khi lưu file CV PDF"})
			}
			student.CvUrl = fmt.Sprintf("http://localhost:3000/uploads/cvs/%s", cvName)
		}

		// Lưu toàn bộ cập nhật vào Database
		if err := db.Save(&student).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Không thể cập nhật DB"})
		}

		return c.JSON(fiber.Map{
			"message": "🎉 Cập nhật hồ sơ sinh viên thành công!",
			"data":    student,
		})
	}
}
