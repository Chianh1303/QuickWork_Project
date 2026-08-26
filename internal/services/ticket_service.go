package services

import (
	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"
	"errors"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type TicketService interface {
	CreateTicket(userID uint, req dto.CreateTicketRequest) (*models.Ticket, error)
	GetUserTickets(userID uint) ([]models.Ticket, error)
	ReappealTicket(userID uint, ticketID uint, req dto.ReappealTicketRequest) error
}

type ticketService struct {
	ticketRepo   repositories.TicketRepository
	appRepo      repositories.ApplicationRepository
	notifService NotificationService
}

func NewTicketService(
	ticketRepo repositories.TicketRepository,
	appRepo repositories.ApplicationRepository,
	notifService NotificationService,
) TicketService {
	return &ticketService{
		ticketRepo:   ticketRepo,
		appRepo:      appRepo,
		notifService: notifService,
	}
}

func (s *ticketService) CreateTicket(userID uint, req dto.CreateTicketRequest) (*models.Ticket, error) {
	if req.ApplicationID == 0 {
		return nil, errors.New("Thiếu mã đơn ứng tuyển (application_id)")
	}

	reasonClean := strings.TrimSpace(req.Reason)
	reasonRunes := []rune(reasonClean)
	if len(reasonRunes) < 3 || len(reasonRunes) > 100 {
		return nil, errors.New("Lý do khiếu nại phải từ 3 đến 100 ký tự")
	}

	descClean := strings.TrimSpace(req.Description)
	descRunes := []rune(descClean)
	if len(descRunes) < 10 || len(descRunes) > 2000 {
		return nil, errors.New("Mô tả chi tiết khiếu nại phải từ 10 đến 2000 ký tự")
	}

	app, err := s.appRepo.GetApplicationByID(req.ApplicationID)
	if err != nil {
		return nil, errors.New("Không tìm thấy đơn ứng tuyển để gửi khiếu nại")
	}

	// Target ID calculation & ownership verification
	var targetUserID uint
	if app.Student.UserID == userID {
		// Reporter is Student, Target is Business User
		job, errJob := s.appRepo.GetJobByID(app.JobID)
		if errJob == nil && job != nil && job.Business.UserID != 0 {
			targetUserID = job.Business.UserID
		}
	} else {
		// Check if Reporter is Business User
		job, errJob := s.appRepo.GetJobByID(app.JobID)
		if errJob == nil && job != nil && job.Business.UserID == userID {
			targetUserID = app.Student.UserID
		}
	}

	if targetUserID == 0 {
		return nil, errors.New("Bạn không có quyền gửi khiếu nại cho đơn ứng tuyển này")
	}

	if targetUserID == userID {
		return nil, errors.New("Bạn không thể tự khiếu nại chính mình")
	}

	// Application-level duplicate check
	exist, _ := s.ticketRepo.GetTicketByApplicationAndReporter(req.ApplicationID, userID)
	if exist != nil {
		return nil, errors.New("Bạn đã gửi khiếu nại cho đơn ứng tuyển này rồi. Vui lòng chờ Ban quản trị xử lý!")
	}

	reqActionClean := strings.TrimSpace(req.RequestedAction)
	evidenceClean := strings.TrimSpace(req.EvidenceURL)

	ticket := &models.Ticket{
		ApplicationID:   req.ApplicationID,
		ReporterID:      userID,
		TargetID:        targetUserID,
		Reason:          reasonClean,
		Description:     descClean,
		RequestedAction: reqActionClean,
		EvidenceURL:     evidenceClean,
		Status:          "pending",
	}

	if err := s.ticketRepo.CreateTicket(ticket); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "1062") || strings.Contains(err.Error(), "idx_ticket_app_reporter") {
			return nil, errors.New("Bạn đã gửi khiếu nại cho đơn ứng tuyển này rồi. Vui lòng chờ Ban quản trị xử lý!")
		}
		return nil, err
	}

	// Send notification to Target User (Side effect isolation)
	if s.notifService != nil {
		_ = s.notifService.CreateNotification(
			targetUserID,
			"⚠️ Thông báo khiếu nại đơn ứng tuyển",
			fmt.Sprintf("Đơn ứng tuyển #%d vừa có khiếu nại từ đối phương. Ban quản trị sẽ đối soát và giải quyết.", app.ID),
			"ticket",
			ticket.ID,
		)
		log.Printf("🔔 [CreateTicket Notification]: Sent to Target UserID=%d for TicketID=%d", targetUserID, ticket.ID)
	}

	return ticket, nil
}

func (s *ticketService) GetUserTickets(userID uint) ([]models.Ticket, error) {
	return s.ticketRepo.GetUserTickets(userID)
}

func (s *ticketService) ReappealTicket(userID uint, ticketID uint, req dto.ReappealTicketRequest) error {
	if ticketID < 1 {
		return errors.New("Mã khiếu nại không hợp lệ")
	}
	reasonClean := strings.TrimSpace(req.Reason)
	runes := []rune(reasonClean)
	if len(runes) < 10 || len(runes) > 1000 {
		return errors.New("Lý do yêu cầu tái xem xét phán quyết phải từ 10 đến 1000 ký tự")
	}

	evidenceClean := strings.TrimSpace(req.Evidence)
	rowsAffected, err := s.ticketRepo.ReappealTicket(ticketID, userID, reasonClean, evidenceClean)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Không thể yêu cầu tái xem xét (khiếu nại không tồn tại, bạn không có quyền hoặc khiếu nại chưa được phán quyết)")
	}

	return nil
}
