package services

import (
	"errors"

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
	StudentCompleteJob(userID uint, appID uint) (*models.Application, error)
	GetEmployerApplications(userID uint) ([]models.Application, error)
	ReviewApplication(userID uint, input dto.ReviewApplicationInput) (*models.Application, string, error)
	BusinessCompleteJob(userID uint, appID uint) (*models.Application, error)
	RespondToOffer(userID uint, input dto.RespondOfferInput) (*models.Application, string, error)
}

type applicationService struct {
	appRepo repositories.ApplicationRepository
}

func NewApplicationService(appRepo repositories.ApplicationRepository) ApplicationService {
	return &applicationService{appRepo: appRepo}
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

func (s *applicationService) StudentCompleteJob(userID uint, appID uint) (*models.Application, error) {
	student, err := s.appRepo.GetStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentAppProfileNotFound
		}
		return nil, err
	}

	app, err := s.appRepo.GetApplicationByIDAndStudent(appID, student.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppNotFound
		}
		return nil, err
	}

	if app.Status != "offer_accepted" {
		return nil, ErrMustBeOfferAcceptedToComplete
	}

	app.Status = "student_completed"
	if err := s.appRepo.SaveApplication(app); err != nil {
		return nil, err
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

	if input.Status != "approved" && input.Status != "rejected" {
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

	msg := "🎉 Bạn đã đồng ý nhận offer công việc thành công!"
	if input.Response == "decline" {
		msg = "❌ Bạn đã từ chối offer công việc thành công."
	}

	return app, msg, nil
}
