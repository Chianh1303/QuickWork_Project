package controllers

import (
	"errors"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AttendanceController struct {
	attendanceService services.AttendanceService
}

func NewAttendanceController(attendanceService services.AttendanceService) *AttendanceController {
	return &AttendanceController{attendanceService: attendanceService}
}

// CheckIn POST /api/attendance/check-in
func (ctrl *AttendanceController) CheckIn(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Hết phiên đăng nhập, vui lòng đăng nhập lại!"})
	}
	userID, ok := userIDVal.(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dữ liệu User ID không hợp lệ"})
	}

	var input dto.AttendanceInput
	if err := c.BodyParser(&input); err != nil || input.JobID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Thiếu mã công việc (Job ID)"})
	}

	attendance, err := ctrl.attendanceService.CheckIn(uint(userID), input.JobID)
	if err != nil {
		if errors.Is(err, services.ErrAttendanceStudentNotFoundCheckIn) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrActiveShiftExists) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, services.ErrNotAppliedToJob) || errors.Is(err, services.ErrOfferNotAcceptedForShift) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Lỗi hệ thống, không thể lưu dữ liệu chấm công"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "⚡ Check-in đầu ca thành công! Chúc bạn làm việc vui vẻ.",
		"data":    attendance,
	})
}

// CheckOut POST /api/attendance/check-out
func (ctrl *AttendanceController) CheckOut(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Hết phiên đăng nhập!"})
	}
	userID, ok := userIDVal.(float64)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dữ liệu User ID không hợp lệ"})
	}

	var input dto.AttendanceInput
	if err := c.BodyParser(&input); err != nil || input.JobID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Thiếu mã công việc (Job ID)"})
	}

	attendance, err := ctrl.attendanceService.CheckOut(uint(userID), input.JobID)
	if err != nil {
		if errors.Is(err, services.ErrAttendanceStudentNotFoundCheckOut) || errors.Is(err, services.ErrNoOpenShiftToCheckout) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Không thể cập nhật thông tin ra ca"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "🛑 Check-out ra ca thành công! Hệ thống đã ghi nhận ca làm của bạn.",
		"data":    attendance,
	})
}
