package controllers

import (
	"errors"
	"strconv"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

// GetAdminDashboardStats GET /api/admin/dashboard/stats
func (ctrl *AdminController) GetAdminDashboardStats(c *fiber.Ctx) error {
	stats, err := ctrl.adminService.GetDashboardStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
	}
	return c.JSON(stats)
}

// GetPendingBusinesses GET /api/admin/businesses/pending
func (ctrl *AdminController) GetPendingBusinesses(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "page must be greater than or equal to 1",
		})
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "limit must be between 1 and 100",
		})
	}

	search := c.Query("search")
	res, err := ctrl.adminService.GetPendingBusinesses(page, limit, search)
	if err != nil {
		if errors.Is(err, services.ErrPageInvalid) || errors.Is(err, services.ErrLimitInvalid) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to count pending businesses",
		})
	}

	return c.JSON(res)
}

// GetBusinessKYBDetail GET /api/admin/businesses/:id
func (ctrl *AdminController) GetBusinessKYBDetail(c *fiber.Ctx) error {
	businessID, err := strconv.Atoi(c.Params("id"))
	if err != nil || businessID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "business id is invalid"})
	}

	detail, err := ctrl.adminService.GetBusinessKYBDetail(businessID)
	if err != nil {
		if errors.Is(err, services.ErrBusinessIDInvalid) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrAdminBusinessNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get business detail"})
	}

	return c.JSON(detail)
}

// ReviewBusinessKYB PUT /api/admin/businesses/:id/review
func (ctrl *AdminController) ReviewBusinessKYB(c *fiber.Ctx) error {
	businessID, err := strconv.Atoi(c.Params("id"))
	if err != nil || businessID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "business id is invalid"})
	}

	adminID, ok := localUserID(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token không hợp lệ"})
	}

	var req dto.ReviewBusinessRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	err = ctrl.adminService.ReviewBusinessKYB(businessID, adminID, req)
	if err != nil {
		if errors.Is(err, services.ErrBusinessIDInvalid) || errors.Is(err, services.ErrDecisionInvalid) || errors.Is(err, services.ErrRejectReasonTooShort) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrProfileAlreadyHandled) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrAdminBusinessNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to review business"})
	}

	return c.JSON(fiber.Map{
		"message":  "Đã cập nhật kết quả KYB doanh nghiệp",
		"decision": req.Decision,
	})
}

func localUserID(c *fiber.Ctx) (uint, bool) {
	switch value := c.Locals("user_id").(type) {
	case float64:
		return uint(value), true
	case uint:
		return value, true
	case int:
		return uint(value), true
	default:
		return 0, false
	}
}
