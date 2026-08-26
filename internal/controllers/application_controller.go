package controllers

import (
	"errors"
	"fmt"
	"log"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ApplicationController struct {
	appService services.ApplicationService
}

func NewApplicationController(appService services.ApplicationService) *ApplicationController {
	return &ApplicationController{appService: appService}
}

// GetStudentApplications GET /api/applications/my-applications
func (ctrl *ApplicationController) GetStudentApplications(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	apps, err := ctrl.appService.GetStudentApplications(userID)
	if err != nil {
		if errors.Is(err, services.ErrStudentAppProfileNotFoundCustom) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "❌ Lỗi hệ thống khi lấy lịch sử ứng tuyển"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    apps,
	})
}

// ApplyJob POST /api/jobs/apply
func (ctrl *ApplicationController) ApplyJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.ApplyJobInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
	}

	app, err := ctrl.appService.ApplyJob(userID, input)
	if err != nil {
		if errors.Is(err, services.ErrStudentAppProfileNotFound) || errors.Is(err, services.ErrJobNotFoundForApp) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrAlreadyApplied) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể nộp đơn ứng tuyển lúc này"})
	}

	return c.JSON(fiber.Map{
		"message": "🚀 Nộp đơn ứng tuyển thành công! Đang chờ Doanh nghiệp phản hồi.",
		"data":    app,
	})
}

// CancelApplication POST /api/applications/:id/cancel
func (ctrl *ApplicationController) CancelApplication(c *fiber.Ctx) error {
	appID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID đơn ứng tuyển không hợp lệ"})
	}

	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	if err := ctrl.appService.CancelApplication(userID, appID); err != nil {
		if errors.Is(err, services.ErrStudentAppProfileNotFound) || errors.Is(err, services.ErrAppNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrCannotCancelProcessedApp) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể hủy đơn lúc này, vui lòng thử lại"})
	}

	return c.JSON(fiber.Map{
		"message": "❌ Đã hủy đơn ứng tuyển thành công!",
	})
}

// StudentCompleteJob POST /api/applications/student-complete
func (ctrl *ApplicationController) StudentCompleteJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.CompleteJobInput
	if err := c.BodyParser(&input); err != nil || input.ApplicationID == 0 {
		log.Printf("❌ [StudentCompleteJob Input Error]: BodyParser err=%v, app_id=%d, Body: %s", err, input.ApplicationID, string(c.Body()))
		return c.Status(400).JSON(fiber.Map{"error": "Thiếu application_id không hợp lệ"})
	}

	app, err := ctrl.appService.StudentCompleteJob(userID, input)
	if err != nil {
		log.Printf("❌ [StudentCompleteJob Service Error]: %v | UserID=%d, AppID=%d", err, userID, input.ApplicationID)
		if errors.Is(err, services.ErrStudentAppProfileNotFound) || errors.Is(err, services.ErrAppNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrMustBeOfferAcceptedToComplete) {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Không thể cập nhật trạng thái"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Sinh viên đã xác nhận hoàn thành công việc. Đang chờ doanh nghiệp xác nhận.",
		"data":    app,
	})
}

// GetEmployerApplications GET /api/applications/employer
func (ctrl *ApplicationController) GetEmployerApplications(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	apps, err := ctrl.appService.GetEmployerApplications(userID)
	if err != nil {
		if errors.Is(err, services.ErrEmployerAppProfileNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "❌ Lỗi hệ thống khi quét danh sách ứng viên"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    apps,
	})
}

// ReviewApplication PUT /api/jobs/review-application
func (ctrl *ApplicationController) ReviewApplication(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.ReviewApplicationInput
	if err := c.BodyParser(&input); err != nil {
		log.Printf("❌ [ReviewApplication BodyParser Error]: %v | Body: %s", err, string(c.Body()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": fmt.Sprintf("Dữ liệu không hợp lệ: %v", err)})
	}

	if input.ApplicationID == 0 {
		log.Printf("❌ [ReviewApplication Input Error]: application_id is 0 | Body: %s", string(c.Body()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "application_id không được để trống hoặc bằng 0"})
	}

	log.Printf("📥 [ReviewApplication Request]: UserID=%d, AppID=%d, Status='%s', OfferSalary='%s'", userID, input.ApplicationID, input.Status, input.OfferSalary)

	app, msg, err := ctrl.appService.ReviewApplication(userID, input)
	if err != nil {
		log.Printf("❌ [ReviewApplication Service Error]: %v", err)
		if errors.Is(err, services.ErrBusinessAppProfileNotFound) || errors.Is(err, services.ErrAppNotFoundForReview) || errors.Is(err, services.ErrJobNotFoundForReview) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrInvalidStatus) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrUnauthorizedAppReview) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Lỗi hệ thống khi cập nhật phản hồi"})
	}

	return c.JSON(fiber.Map{
		"message": msg,
		"data":    app,
	})
}

// BusinessCompleteJob POST /api/applications/business-complete
func (ctrl *ApplicationController) BusinessCompleteJob(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.CompleteJobInput
	if err := c.BodyParser(&input); err != nil || input.ApplicationID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Thiếu application_id"})
	}

	app, err := ctrl.appService.BusinessCompleteJob(userID, input.ApplicationID)
	if err != nil {
		if errors.Is(err, services.ErrBusinessAppProfileNotFound) || errors.Is(err, services.ErrAppNotFoundForReview) {
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrForbiddenConfirmBusiness) {
			return c.Status(403).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrStudentNotCompletedYet) {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Không thể xác nhận hoàn thành"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Doanh nghiệp đã xác nhận thanh toán thành công!",
		"data":    app,
	})
}

// RespondToOffer POST /api/offers/respond
func (ctrl *ApplicationController) RespondToOffer(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Chưa đăng nhập"})
	}
	userID := uint(userIDVal.(float64))

	var input dto.RespondOfferInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "❌ Dữ liệu phản hồi không hợp lệ"})
	}

	app, msg, err := ctrl.appService.RespondToOffer(userID, input)
	if err != nil {
		if errors.Is(err, services.ErrStudentAppProfileNotFound) || errors.Is(err, services.ErrOfferAppNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if errors.Is(err, services.ErrOfferAppNoActiveOffer) || errors.Is(err, services.ErrOfferInvalidAction) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "❌ Lỗi hệ thống khi lưu phản hồi"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": msg,
		"data":    app,
	})
}
