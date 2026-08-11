package controllers

import (
	"errors"
	"strconv"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ReviewController struct {
	reviewService services.ReviewService
}

func NewReviewController(reviewService services.ReviewService) *ReviewController {
	return &ReviewController{reviewService: reviewService}
}

// GetReviewsByUser GET /api/reviews/user/:userId
func (ctrl *ReviewController) GetReviewsByUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("userId"), 10, 64)
	if err != nil || userID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userId không hợp lệ",
		})
	}

	reviews, average, total, err := ctrl.reviewService.GetReviewsByUser(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể tải danh sách đánh giá",
		})
	}

	return c.JSON(fiber.Map{
		"success":        true,
		"average_rating": average,
		"total_reviews":  total,
		"reviews":        reviews,
	})
}

// GetReviewsByApplication GET /api/reviews/application/:applicationId
func (ctrl *ReviewController) GetReviewsByApplication(c *fiber.Ctx) error {
	applicationID, err := strconv.ParseUint(c.Params("applicationId"), 10, 64)
	if err != nil || applicationID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "applicationId không hợp lệ",
		})
	}

	reviews, err := ctrl.reviewService.GetReviewsByApplication(uint(applicationID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể tải danh sách đánh giá",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"reviews": reviews,
	})
}

// CreateReview POST /api/reviews
func (ctrl *ReviewController) CreateReview(c *fiber.Ctx) error {
	userIDVal, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token không hợp lệ",
		})
	}

	var input dto.CreateReviewInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Body không hợp lệ",
		})
	}

	review, err := ctrl.reviewService.CreateReview(uint(userIDVal), input)
	if err != nil {
		if errors.Is(err, services.ErrMissingApplicationID) || errors.Is(err, services.ErrRatingOutOfBounds) || errors.Is(err, services.ErrJobNotFinished) || errors.Is(err, services.ErrAlreadyReviewed) {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrReviewAppNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrNoPermissionToReview) {
			return c.Status(403).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Không thể lưu Review"})
	}

	return c.JSON(fiber.Map{
		"message": "Đánh giá thành công.",
		"data":    review,
	})
}
