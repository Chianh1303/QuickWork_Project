package controllers

import (
	"io"

	"QuickWork/internal/dto"
	"QuickWork/internal/services"
	"QuickWork/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type TicketController struct {
	ticketService   services.TicketService
	storageProvider storage.StorageProvider
}

func NewTicketController(ticketService services.TicketService, storageProvider ...storage.StorageProvider) *TicketController {
	var sp storage.StorageProvider
	if len(storageProvider) > 0 && storageProvider[0] != nil {
		sp = storageProvider[0]
	} else {
		sp = storage.NewStorageProvider()
	}
	return &TicketController{
		ticketService:   ticketService,
		storageProvider: sp,
	}
}

// UploadEvidence POST /api/tickets/upload-evidence
func (ctrl *TicketController) UploadEvidence(c *fiber.Ctx) error {
	_, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập hoặc phiên làm việc không hợp lệ",
		})
	}

	file, err := c.FormFile("evidence")
	if err != nil || file == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Vui lòng chọn file hình ảnh/bằng chứng để tải lên",
		})
	}

	if file.Size > 10*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Dung lượng file bằng chứng không được vượt quá 10MB",
		})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể mở file bằng chứng",
		})
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể đọc dữ liệu file bằng chứng",
		})
	}

	url, err := ctrl.storageProvider.UploadFile(fileBytes, file.Filename, "disputes")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Tải file bằng chứng thất bại: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tải file bằng chứng thành công!",
		"url":     url,
	})
}

// CreateTicket POST /api/tickets
func (ctrl *TicketController) CreateTicket(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập hoặc phiên làm việc không hợp lệ",
		})
	}

	var req dto.CreateTicketRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Dữ liệu khiếu nại không đúng định dạng JSON",
		})
	}

	ticket, err := ctrl.ticketService.CreateTicket(userID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "🎉 Gửi khiếu nại thành công! Ban quản trị QuickWork sẽ tiến hành đối soát và xử lý.",
		"ticket":  ticket,
	})
}

// GetUserTickets GET /api/tickets/my-tickets
func (ctrl *TicketController) GetUserTickets(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập hoặc phiên làm việc không hợp lệ",
		})
	}

	tickets, err := ctrl.ticketService.GetUserTickets(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Không thể lấy danh sách khiếu nại",
		})
	}

	return c.JSON(fiber.Map{
		"data": tickets,
	})
}

// ReappealTicket POST /api/tickets/:id/reappeal
func (ctrl *TicketController) ReappealTicket(c *fiber.Ctx) error {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Chưa đăng nhập hoặc phiên làm việc không hợp lệ",
		})
	}

	ticketID, err := c.ParamsInt("id")
	if err != nil || ticketID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Mã khiếu nại không hợp lệ",
		})
	}

	var req dto.ReappealTicketRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Dữ liệu yêu cầu không đúng định dạng JSON",
		})
	}

	if err := ctrl.ticketService.ReappealTicket(userID, uint(ticketID), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "🎉 Đã gửi yêu cầu tái xem xét phán quyết thành công! Ban quản trị sẽ tiến hành đối soát lại cấp cao hơn.",
	})
}
