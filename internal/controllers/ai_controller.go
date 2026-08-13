package controllers

import (
	"errors"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AIController struct {
	aiService services.AIService
}

func NewAIController(aiService services.AIService) *AIController {
	return &AIController{aiService: aiService}
}

// EvaluateCV POST /api/ai/evaluate-cv
func (ctrl *AIController) EvaluateCV(c *fiber.Ctx) error {
	var req dto.EvaluateCVRequest
	if len(c.Body()) > 0 {
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Dữ liệu request không hợp lệ",
			})
		}
	}

	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập",
		})
	}

	res, err := ctrl.aiService.EvaluateCV(userID, req)
	if err != nil {
		if errors.Is(err, services.ErrStudentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể lấy hồ sơ sinh viên",
		})
	}

	return c.JSON(res)
}

// MatchJob POST /api/ai/match-job
func (ctrl *AIController) MatchJob(c *fiber.Ctx) error {
	var req dto.MatchJobRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Dữ liệu không hợp lệ",
		})
	}

	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập",
		})
	}

	res, err := ctrl.aiService.MatchJob(userID, req)
	if err != nil {
		if errors.Is(err, services.ErrStudentNotFound) || errors.Is(err, services.ErrJobNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể thực hiện đánh giá độ phù hợp",
		})
	}

	return c.JSON(res)
}

// GenerateJobDescription POST /api/ai/generate-job
func (ctrl *AIController) GenerateJobDescription(c *fiber.Ctx) error {
	var req dto.GenerateJobRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Vui lòng nhập tên vị trí tuyển dụng",
		})
	}

	res, err := ctrl.aiService.GenerateJobDescription(req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidInput) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Vui lòng nhập tên vị trí tuyển dụng",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể tạo mô tả tin tuyển dụng",
		})
	}

	return c.JSON(res)
}

// GetRecommendedJobs GET /api/ai/recommended-jobs
func (ctrl *AIController) GetRecommendedJobs(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập",
		})
	}

	res, err := ctrl.aiService.GetRecommendedJobs(userID)
	if err != nil {
		if errors.Is(err, services.ErrStudentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể lấy danh sách công việc gợi ý",
		})
	}

	return c.JSON(res)
}

func getUserIDFromContext(c *fiber.Ctx) (uint, bool) {
	val := c.Locals("user_id")
	if val == nil {
		val = c.Locals("userID")
	}
	if val == nil {
		return 0, false
	}
	switch v := val.(type) {
	case float64:
		return uint(v), true
	case uint:
		return v, true
	case int:
		return uint(v), true
	case int64:
		return uint(v), true
	default:
		return 0, false
	}
}
