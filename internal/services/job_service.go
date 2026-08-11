package services

import (
	"errors"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrBusinessProfileMissing = errors.New("Không tìm thấy hồ sơ doanh nghiệp của tài khoản này")
	ErrUserAccountNotFound    = errors.New("Không tìm thấy tài khoản")
	ErrBusinessNotApproved    = errors.New("Tài khoản doanh nghiệp của bạn đang chờ duyệt hoặc đã bị từ chối.")
	ErrJobInputInvalid        = errors.New("Vui lòng điền đầy đủ tiêu đề, mô tả, địa điểm và lương lớn hơn 0")
)

type JobService interface {
	GetAvailableJobs(search, location, category, jobType, maxSalary string) ([]models.Job, error)
	CreateJob(userID uint, input dto.CreateJobInput) (*models.Job, error)
}

type jobService struct {
	jobRepo repositories.JobRepository
}

func NewJobService(jobRepo repositories.JobRepository) JobService {
	return &jobService{jobRepo: jobRepo}
}

func (s *jobService) GetAvailableJobs(search, location, category, jobType, maxSalary string) ([]models.Job, error) {
	return s.jobRepo.GetAvailableJobs(search, location, category, jobType, maxSalary)
}

func (s *jobService) CreateJob(userID uint, input dto.CreateJobInput) (*models.Job, error) {
	business, user, err := s.jobRepo.GetBusinessAndUserByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if business == nil {
				return nil, ErrBusinessProfileMissing
			}
			return nil, ErrUserAccountNotFound
		}
		return nil, err
	}

	if user.Status != "approved" && user.Status != "active" {
		return nil, ErrBusinessNotApproved
	}

	if input.Title == "" || input.Description == "" || input.Location == "" || input.Salary <= 0 {
		return nil, ErrJobInputInvalid
	}

	newJob := &models.Job{
		BusinessID:  business.ID,
		Title:       input.Title,
		Description: input.Description,
		Location:    input.Location,
		Salary:      input.Salary,
		Slots:       input.Slots,
		Status:      "pending",
		WorkingDate: input.WorkingDate,
	}

	if err := s.jobRepo.CreateJob(newJob); err != nil {
		return nil, err
	}

	return newJob, nil
}
