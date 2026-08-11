package controllers

import (
	"errors"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type JobController struct {
	jobService services.JobService
}

func NewJobController(jobService services.JobService) *JobController {
	return &JobController{jobService: jobService}
}

// GetAvailableJobs GET /api/jobs
func (ctrl *JobController) GetAvailableJobs(c *fiber.Ctx) error {
	search := c.Query("search")
	location := c.Query("location")
	category := c.Query("category")
	jobType := c.Query("job_type")
	maxSalary := c.Query("max_salary")

	jobs, err := ctrl.jobService.GetAvailableJobs(search, location, category, jobType, maxSalary)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể lấy danh sách công việc toàn diện",
		})
	}

	return c.JSON(fiber.Map{
		"message": "🎨 Lấy danh sách công việc thành công!",
		"data":    jobs,
	})
}

// CreateJob POST /api/jobs
func (ctrl *JobController) CreateJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.CreateJobInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	job, err := ctrl.jobService.CreateJob(userID, input)
	if err != nil {
		if errors.Is(err, services.ErrBusinessProfileMissing) || errors.Is(err, services.ErrUserAccountNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrBusinessNotApproved) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrJobInputInvalid) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể đăng tin tuyển dụng lúc này"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Đăng tin thành công! Vui lòng chờ Admin phê duyệt",
		"data":    job,
	})
}
