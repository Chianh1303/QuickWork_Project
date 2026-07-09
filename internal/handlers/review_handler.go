package handlers

import (
	"QuickWork/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetReviewsByUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := strconv.ParseUint(c.Params("userId"), 10, 64)
		if err != nil || userID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "userId không hợp lệ",
			})
		}

		var reviews []models.Review
		if err := db.
			Where("reviewee_id = ?", userID).
			Order("created_at DESC").
			Find(&reviews).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Không thể tải danh sách đánh giá",
			})
		}

		total := len(reviews)
		var average float64
		if total > 0 {
			var sum int
			for _, review := range reviews {
				sum += review.Rating
			}
			average = float64(sum) / float64(total)
		}

		return c.JSON(fiber.Map{
			"success":        true,
			"average_rating": average,
			"total_reviews":  total,
			"reviews":        reviews,
		})
	}
}

func GetReviewsByApplication(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		applicationID, err := strconv.ParseUint(c.Params("applicationId"), 10, 64)
		if err != nil || applicationID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "applicationId không hợp lệ",
			})
		}

		var reviews []models.Review
		if err := db.
			Where("application_id = ?", applicationID).
			Order("created_at DESC").
			Find(&reviews).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Không thể tải danh sách đánh giá",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"reviews": reviews,
		})
	}
}

func CreateReview(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userID, ok := c.Locals("user_id").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token không hợp lệ",
			})
		}

		var input struct {
			ApplicationID uint   `json:"application_id"`
			Rating        int    `json:"rating"`
			Comment       string `json:"comment"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body không hợp lệ",
			})
		}

		if input.ApplicationID == 0 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Thiếu application_id",
			})
		}

		if input.Rating < 1 || input.Rating > 5 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Rating phải từ 1 đến 5",
			})
		}
		var app models.Application

		if err := db.
			Preload("Job.Business").
			Preload("Student").
			First(&app, input.ApplicationID).Error; err != nil {

			return c.Status(404).JSON(fiber.Map{
				"error": "Không tìm thấy Application",
			})
		}
		if app.Status != "paid" {

			return c.Status(400).JSON(fiber.Map{
				"error": "Công việc chưa hoàn thành.",
			})

		}

		var reviewerID uint
		var revieweeID uint
		var reviewType string

		if app.Student.UserID == uint(userID) {

			reviewerID = app.Student.UserID

			revieweeID = app.Job.Business.UserID

			reviewType = "business"

		}
		if app.Job.Business.UserID == uint(userID) {

			reviewerID = app.Job.Business.UserID

			revieweeID = app.Student.UserID

			reviewType = "student"

		}
		if reviewerID == 0 {

			return c.Status(403).JSON(fiber.Map{
				"error": "Bạn không có quyền đánh giá.",
			})

		}
		var existing models.Review

		err := db.Where(
			"application_id=? AND reviewer_id=?",
			input.ApplicationID,
			reviewerID,
		).First(&existing).Error

		if err == nil {

			return c.Status(400).JSON(fiber.Map{
				"error": "Bạn đã đánh giá rồi.",
			})

		}
		review := models.Review{

			ApplicationID: input.ApplicationID,

			ReviewerID: reviewerID,

			RevieweeID: revieweeID,

			ReviewType: reviewType,

			Rating: input.Rating,

			Comment: input.Comment,
		}
		if err := db.Create(&review).Error; err != nil {

			return c.Status(500).JSON(fiber.Map{
				"error": "Không thể lưu Review",
			})

		}
		return c.JSON(fiber.Map{

			"message": "Đánh giá thành công.",

			"data": review,
		})
	}
}
