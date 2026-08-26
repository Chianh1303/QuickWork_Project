package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrStudentAppProfileNotFound = errors.New("Không tìm thấy hồ sơ sinh viên")
	ErrStudentAppProfileNotFoundCustom = errors.New("❌ Không tìm thấy hồ sơ sinh viên của tài khoản này")
	ErrJobNotFoundForApp         = errors.New("Công việc này không tồn tại")
	ErrAlreadyApplied            = errors.New("⚠️ Bạn đã ứng tuyển công việc này rồi!")
	ErrAppNotFound               = errors.New("Không tìm thấy đơn ứng tuyển này của bạn")
	ErrCannotCancelProcessedApp  = errors.New("⚠️ Không thể hủy đơn ứng tuyển đã được doanh nghiệp xử lý!")
	ErrMustBeOfferAcceptedToComplete = errors.New("Chỉ có thể hoàn thành khi đã nhận việc")
	ErrBusinessAppProfileNotFound = errors.New("Không tìm thấy tài khoản Doanh nghiệp")
	ErrEmployerAppProfileNotFound = errors.New("❌ Không tìm thấy thông tin doanh nghiệp")
	ErrInvalidStatus             = errors.New("Trạng thái không hợp lệ")
	ErrAppNotFoundForReview      = errors.New("Không tìm thấy đơn ứng tuyển này")
	ErrJobNotFoundForReview      = errors.New("Không tìm thấy tin tuyển dụng này")
	ErrUnauthorizedAppReview     = errors.New("Bạn không có quyền chỉnh sửa đơn ứng tuyển này")
	ErrStudentNotCompletedYet    = errors.New("Sinh viên chưa xác nhận hoàn thành")
	ErrForbiddenConfirmBusiness  = errors.New("Bạn không có quyền xác nhận đơn này")
	ErrOfferAppNotFound          = errors.New("❌ Không tìm thấy đơn ứng tuyển tương ứng")
	ErrOfferAppNoActiveOffer     = errors.New("❌ Đơn ứng tuyển này hiện không có Offer nào cần xử lý")
	ErrOfferInvalidAction        = errors.New("❌ Hành động không hợp lệ")
)

type ApplicationService interface {
	GetStudentApplications(userID uint) ([]models.Application, error)
	ApplyJob(userID uint, input dto.ApplyJobInput) (*models.Application, error)
	CancelApplication(userID uint, appID int) error
	StudentCompleteJob(userID uint, input dto.CompleteJobInput) (*models.Application, error)
	GetEmployerApplications(userID uint) ([]models.Application, error)
	ReviewApplication(userID uint, input dto.ReviewApplicationInput) (*models.Application, string, error)
	BusinessCompleteJob(userID uint, appID uint) (*models.Application, error)
	RespondToOffer(userID uint, input dto.RespondOfferInput) (*models.Application, string, error)
}

type applicationService struct {
	appRepo      repositories.ApplicationRepository
	notifService NotificationService
}

func NewApplicationService(appRepo repositories.ApplicationRepository, notifService ...NotificationService) ApplicationService {
	var ns NotificationService
	if len(notifService) > 0 {
		ns = notifService[0]
	}
	return &applicationService{
		appRepo:      appRepo,
		notifService: ns,
	}
}

func (s *applicationService) GetStudentApplications(userID uint) ([]models.Application, error) {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentAppProfileNotFoundCustom
		}
		return nil, err
	}
	return s.appRepo.GetStudentApplications(student.ID)
}

func (s *applicationService) ApplyJob(userID uint, input dto.ApplyJobInput) (*models.Application, error) {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentAppProfileNotFound
		}
		return nil, err
	}

	_, err = s.appRepo.GetJobByID(input.JobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrJobNotFoundForApp
		}
		return nil, err
	}

	existApp, err := s.appRepo.GetApplicationByJobAndStudent(input.JobID, student.ID)
	if err == nil && existApp != nil {
		if existApp.Status == "rejected" || existApp.Status == "offer_declined" {
			// 🌟 NÂNG CẤP CHUẨN NGHIỆP VỤ: Cho phép ứng viên nộp lại đơn nếu trước đó bị từ chối
			existApp.Status = "pending"
			existApp.CoverNote = input.CoverNote
			existApp.OfferSalary = ""
			existApp.OfferStartDate = ""
			existApp.OfferMessage = ""
			existApp.StudentCompleted = false
			existApp.CompletionNote = ""
			existApp.CompletionProofUrl = ""

			if errSave := s.appRepo.SaveApplication(existApp); errSave != nil {
				return nil, errSave
			}

			if s.notifService != nil {
				job, _ := s.appRepo.GetJobByID(input.JobID)
				if job != nil && job.Business.UserID != 0 {
					errNotif := s.notifService.CreateNotification(
						job.Business.UserID,
						"📄 Đơn ứng tuyển mới (Nộp lại)",
						fmt.Sprintf("Ứng viên %s vừa nộp lại đơn ứng tuyển cho vị trí '%s'.", student.FullName, job.Title),
						"application",
						existApp.ID,
					)
					log.Printf("🔔 [Re-ApplyJob Notification]: Gửi thông báo cho Business UserID=%d, error=%v", job.Business.UserID, errNotif)
				}
			}

			return existApp, nil
		}

		return nil, ErrAlreadyApplied
	}

	newApp := &models.Application{
		JobID:     input.JobID,
		StudentID: student.ID,
		Status:    "pending",
		CoverNote: input.CoverNote,
	}

	if err := s.appRepo.CreateApplication(newApp); err != nil {
		return nil, err
	}

	if s.notifService != nil {
		job, _ := s.appRepo.GetJobByID(input.JobID)
		if job != nil && job.Business.UserID != 0 {
			errNotif := s.notifService.CreateNotification(
				job.Business.UserID,
				"📄 Đơn ứng tuyển mới",
				fmt.Sprintf("Ứng viên %s vừa nộp đơn ứng tuyển vị trí '%s'.", student.FullName, job.Title),
				"application",
				newApp.ID,
			)
			log.Printf("🔔 [ApplyJob Notification]: Gửi thông báo cho Business UserID=%d, error=%v", job.Business.UserID, errNotif)
		} else if job != nil {
			log.Printf("⚠️ [ApplyJob Notification Warning]: job.Business.UserID is 0! BusinessID=%d", job.BusinessID)
		}
	}

	return newApp, nil
}

func (s *applicationService) CancelApplication(userID uint, appID int) error {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrStudentAppProfileNotFound
		}
		return err
	}

	app, err := s.appRepo.GetApplicationByIDAndStudent(uint(appID), student.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAppNotFound
		}
		return err
	}

	if app.Status != "pending" {
		return ErrCannotCancelProcessedApp
	}

	return s.appRepo.DeleteApplication(app)
}

func (s *applicationService) StudentCompleteJob(userID uint, input dto.CompleteJobInput) (*models.Application, error) {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentAppProfileNotFound
		}
		return nil, err
	}

	app, err := s.appRepo.GetApplicationByIDAndStudent(input.ApplicationID, student.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppNotFound
		}
		return nil, err
	}

	if app.Status != "offer_accepted" && app.Status != "approved" && app.Status != "student_completed" {
		return nil, ErrMustBeOfferAcceptedToComplete
	}

	now := time.Now()
	app.Status = "student_completed"
	app.StudentCompleted = true
	app.CompletionNote = input.CompletionNote
	app.CompletionProofUrl = input.CompletionProofUrl
	app.SubmittedAt = &now

	if err := s.appRepo.SaveApplication(app); err != nil {
		return nil, err
	}

	if s.notifService != nil {
		job, _ := s.appRepo.GetJobByID(app.JobID)
		if job != nil && job.Business.UserID != 0 {
			errNotif := s.notifService.CreateNotification(
				job.Business.UserID,
				"📝 Báo cáo hoàn thành công việc",
				fmt.Sprintf("Sinh viên %s đã nộp báo cáo hoàn thành cho vị trí '%s'.", student.FullName, job.Title),
				"escrow",
				app.ID,
			)
			log.Printf("🔔 [StudentCompleteJob Notification]: Gửi thông báo cho Business UserID=%d, error=%v", job.Business.UserID, errNotif)
		}
	}

	return app, nil
}

func (s *applicationService) GetEmployerApplications(userID uint) ([]models.Application, error) {
	business, err := s.appRepo.GetBusinessByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEmployerAppProfileNotFound
		}
		return nil, err
	}

	return s.appRepo.GetEmployerApplications(business.ID)
}

func (s *applicationService) ReviewApplication(userID uint, input dto.ReviewApplicationInput) (*models.Application, string, error) {
	business, err := s.appRepo.GetBusinessByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrBusinessAppProfileNotFound
		}
		return nil, "", err
	}

	if input.Status != "approved" && input.Status != "rejected" && input.Status != "pending" {
		return nil, "", ErrInvalidStatus
	}

	app, err := s.appRepo.GetApplicationByID(input.ApplicationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrAppNotFoundForReview
		}
		return nil, "", err
	}

	job, err := s.appRepo.GetJobByID(app.JobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrJobNotFoundForReview
		}
		return nil, "", err
	}

	if job.BusinessID != business.ID {
		return nil, "", ErrUnauthorizedAppReview
	}

	app.Status = input.Status
	if input.Status == "approved" {
		app.OfferSalary = input.OfferSalary
		app.OfferStartDate = input.OfferStartDate
		app.OfferMessage = input.OfferMessage
	}

	if err := s.appRepo.SaveApplication(app); err != nil {
		return nil, "", err
	}

	if s.notifService != nil && (input.Status == "approved" || input.Status == "rejected") {
		var targetUserID uint
		if app.Student.UserID != 0 {
			targetUserID = app.Student.UserID
		} else {
			studentObj, errSt := s.appRepo.GetStudentByID(app.StudentID)
			if errSt == nil && studentObj != nil {
				targetUserID = studentObj.UserID
			}
		}

		if targetUserID != 0 {
			title := "🎉 Bạn nhận được Offer công việc mới!"
			msgText := fmt.Sprintf("Doanh nghiệp %s đã duyệt hồ sơ và gửi kèm Offer công việc cho vị trí '%s'.", business.CompanyName, job.Title)
			if input.Status == "rejected" {
				title = "❌ Cập nhật trạng thái đơn ứng tuyển"
				msgText = fmt.Sprintf("Doanh nghiệp %s đã phản hồi từ chối đơn ứng tuyển cho vị trí '%s'.", business.CompanyName, job.Title)
			}
			errNotif := s.notifService.CreateNotification(
				targetUserID,
				title,
				msgText,
				"offer",
				app.ID,
			)
			log.Printf("🔔 [ReviewApplication Notification Sent]: Target UserID=%d, AppID=%d, Status='%s', err=%v", targetUserID, app.ID, input.Status, errNotif)
		} else {
			log.Printf("⚠️ [ReviewApplication Notification FAILED]: targetUserID is 0 for AppID=%d, StudentID=%d", app.ID, app.StudentID)
		}
	}

	msg := "🎉 Chấp nhận hồ sơ ứng viên và gửi kèm thông tin Offer thành công!"
	if input.Status == "rejected" {
		msg = "❌ Đã từ chối đơn ứng tuyển thành công."
	}

	return app, msg, nil
}

func (s *applicationService) BusinessCompleteJob(userID uint, appID uint) (*models.Application, error) {
	business, err := s.appRepo.GetBusinessByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBusinessAppProfileNotFound
		}
		return nil, err
	}

	app, err := s.appRepo.BusinessCompleteJobAndPay(appID, business.ID)
	if err != nil {
		if err.Error() == "FORBIDDEN_BUSINESS_OWNERSHIP" {
			return nil, ErrForbiddenConfirmBusiness
		}
		if err.Error() == "STUDENT_NOT_COMPLETED" {
			return nil, ErrStudentNotCompletedYet
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppNotFoundForReview
		}
		return nil, err
	}

	if s.notifService != nil && app != nil {
		if app.Student.UserID != 0 {
			jobTitle := "công việc"
			if app.Job.Title != "" {
				jobTitle = app.Job.Title
			}
			_ = s.notifService.CreateNotification(
				app.Student.UserID,
				"💰 Giải ngân tiền lương Escrow",
				fmt.Sprintf("Doanh nghiệp đã xác nhận hoàn thành & giải ngân lương cho vị trí '%s'.", jobTitle),
				"escrow",
				app.ID,
			)
		}
	}

	return app, nil
}

func (s *applicationService) RespondToOffer(userID uint, input dto.RespondOfferInput) (*models.Application, string, error) {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrStudentAppProfileNotFound
		}
		return nil, "", err
	}

	app, err := s.appRepo.GetApplicationByIDAndStudent(input.ApplicationID, student.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrOfferAppNotFound
		}
		return nil, "", err
	}

	if app.Status != "approved" {
		return nil, "", ErrOfferAppNoActiveOffer
	}

	if input.Response == "accept" {
		app.Status = "offer_accepted"
	} else if input.Response == "decline" {
		app.Status = "offer_declined"
	} else {
		return nil, "", ErrOfferInvalidAction
	}

	if err := s.appRepo.SaveApplication(app); err != nil {
		return nil, "", err
	}

	if s.notifService != nil {
		job, _ := s.appRepo.GetJobByID(app.JobID)
		if job != nil && job.Business.UserID != 0 {
			title := "🎉 Ứng viên đã chấp nhận Offer!"
			msgText := fmt.Sprintf("Ứng viên %s đã đồng ý nhận Offer cho vị trí '%s'.", student.FullName, job.Title)
			if input.Response == "decline" {
				title = "❌ Ứng viên đã từ chối Offer"
				msgText = fmt.Sprintf("Ứng viên %s đã từ chối Offer cho vị trí '%s'.", student.FullName, job.Title)
			}
			errNotif := s.notifService.CreateNotification(
				job.Business.UserID,
				title,
				msgText,
				"offer",
				app.ID,
			)
			log.Printf("🔔 [RespondToOffer Notification]: Gửi thông báo cho Business UserID=%d, error=%v", job.Business.UserID, errNotif)
		}
	}

	msg := "🎉 Bạn đã đồng ý nhận offer công việc thành công!"
	if input.Response == "decline" {
		msg = "❌ Bạn đã từ chối offer công việc thành công."
	}

	return app, msg, nil
}
