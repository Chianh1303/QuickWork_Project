package controllers

import (
	"errors"
	"io"
	"path/filepath"

	"QuickWork/internal/services"
	"QuickWork/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	profileService  services.ProfileService
	storageProvider storage.StorageProvider
}

func NewProfileController(profileService services.ProfileService, storageProvider ...storage.StorageProvider) *ProfileController {
	var sp storage.StorageProvider
	if len(storageProvider) > 0 && storageProvider[0] != nil {
		sp = storageProvider[0]
	} else {
		sp = storage.NewStorageProvider()
	}
	return &ProfileController{
		profileService:  profileService,
		storageProvider: sp,
	}
}

// GetStudentProfile GET /api/profile/student
func (ctrl *ProfileController) GetStudentProfile(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	student, err := ctrl.profileService.GetStudentProfile(userID)
	if err != nil {
		if errors.Is(err, services.ErrStudentProfileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    student,
	})
}

// UpdateStudentProfile PUT /api/profile/student
func (ctrl *ProfileController) UpdateStudentProfile(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	fullName := c.FormValue("full_name")
	phone := c.FormValue("phone")
	gender := c.FormValue("gender")
	skills := c.FormValue("skills")

	var avatarUrl string
	avatarFile, err := c.FormFile("avatar")
	if err == nil {
		fileHeader, openErr := avatarFile.Open()
		if openErr == nil {
			fileBytes, readErr := io.ReadAll(fileHeader)
			fileHeader.Close()
			if readErr == nil {
				url, uploadErr := ctrl.storageProvider.UploadFile(fileBytes, avatarFile.Filename, "avatars")
				if uploadErr == nil {
					avatarUrl = url
				}
			}
		}
	}

	var cvUrl string
	cvFile, err := c.FormFile("cv")
	if err == nil {
		if filepath.Ext(cvFile.Filename) != ".pdf" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Hồ sơ đính kèm bắt buộc phải là định dạng file PDF"})
		}
		fileHeader, openErr := cvFile.Open()
		if openErr == nil {
			fileBytes, readErr := io.ReadAll(fileHeader)
			fileHeader.Close()
			if readErr == nil {
				url, uploadErr := ctrl.storageProvider.UploadFile(fileBytes, cvFile.Filename, "cvs")
				if uploadErr == nil {
					cvUrl = url
				}
			}
		}
	}

	student, err := ctrl.profileService.UpdateStudentProfile(userID, fullName, phone, gender, skills, avatarUrl, cvUrl)
	if err != nil {
		if errors.Is(err, services.ErrStudentProfileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Không thể cập nhật DB"})
	}

	return c.JSON(fiber.Map{
		"message": "🎉 Cập nhật hồ sơ sinh viên thành công!",
		"data":    student,
	})
}

// GetBusinessProfile GET /api/profile/business
func (ctrl *ProfileController) GetBusinessProfile(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	business, err := ctrl.profileService.GetBusinessProfile(userID)
	if err != nil {
		if errors.Is(err, services.ErrBusinessProfileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    business,
	})
}

// UpdateBusinessProfile PUT /api/profile/business
func (ctrl *ProfileController) UpdateBusinessProfile(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	companyName := c.FormValue("company_name")
	phone := c.FormValue("phone")
	address := c.FormValue("address")

	var logoUrl string
	logoFile, err := c.FormFile("logo")
	if err == nil {
		fileHeader, openErr := logoFile.Open()
		if openErr == nil {
			fileBytes, readErr := io.ReadAll(fileHeader)
			fileHeader.Close()
			if readErr == nil {
				url, uploadErr := ctrl.storageProvider.UploadFile(fileBytes, logoFile.Filename, "avatars")
				if uploadErr == nil {
					logoUrl = url
				}
			}
		}
	}

	business, err := ctrl.profileService.UpdateBusinessProfile(userID, companyName, phone, address, logoUrl)
	if err != nil {
		if errors.Is(err, services.ErrBusinessProfileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Không thể cập nhật DB doanh nghiệp"})
	}

	return c.JSON(fiber.Map{
		"message": "🎉 Cập nhật hồ sơ doanh nghiệp thành công!",
		"data":    business,
	})
}
