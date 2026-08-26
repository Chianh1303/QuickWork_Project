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
		"message": "Đăng tin thành công!",
		"data":    job,
	})
}

// GetMyBusinessJobs GET /api/jobs/business/my-jobs
func (ctrl *JobController) GetMyBusinessJobs(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	jobs, err := ctrl.jobService.GetMyBusinessJobs(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Lấy danh sách bài tuyển dụng doanh nghiệp thành công",
		"data":    jobs,
	})
}

// UpdateJob PUT /api/jobs/:id
func (ctrl *JobController) UpdateJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	jobID, err := c.ParamsInt("id")
	if err != nil || jobID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID bài tuyển dụng không hợp lệ"})
	}

	var input dto.CreateJobInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	job, err := ctrl.jobService.UpdateJob(userID, uint(jobID), input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Cập nhật bài tuyển dụng thành công!",
		"data":    job,
	})
}

// ToggleJobStatus PATCH /api/jobs/:id/status
func (ctrl *JobController) ToggleJobStatus(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	jobID, err := c.ParamsInt("id")
	if err != nil || jobID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID bài tuyển dụng không hợp lệ"})
	}

	job, err := ctrl.jobService.ToggleJobStatus(userID, uint(jobID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Thay đổi trạng thái bài tuyển dụng thành công",
		"data":    job,
	})
}

// DeleteJob DELETE /api/jobs/:id
func (ctrl *JobController) DeleteJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	jobID, err := c.ParamsInt("id")
	if err != nil || jobID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID bài tuyển dụng không hợp lệ"})
	}

	if err := ctrl.jobService.DeleteJob(userID, uint(jobID)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Đã xóa bài tuyển dụng thành công",
	})
}
