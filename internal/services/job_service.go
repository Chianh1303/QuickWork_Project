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
	GetMyBusinessJobs(userID uint) ([]models.Job, error)
	UpdateJob(userID uint, jobID uint, input dto.CreateJobInput) (*models.Job, error)
	ToggleJobStatus(userID uint, jobID uint) (*models.Job, error)
	DeleteJob(userID uint, jobID uint) error
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
		Status:      "approved", // Business posts are active
		WorkingDate: input.WorkingDate,
		Category:    input.Category,
		JobType:     input.JobType,
	}

	if err := s.jobRepo.CreateJob(newJob); err != nil {
		return nil, err
	}

	return newJob, nil
}

func (s *jobService) GetMyBusinessJobs(userID uint) ([]models.Job, error) {
	business, _, err := s.jobRepo.GetBusinessAndUserByUserID(userID)
	if err != nil {
		return nil, ErrBusinessProfileMissing
	}

	return s.jobRepo.GetJobsByBusinessID(business.ID)
}

func (s *jobService) UpdateJob(userID uint, jobID uint, input dto.CreateJobInput) (*models.Job, error) {
	business, _, err := s.jobRepo.GetBusinessAndUserByUserID(userID)
	if err != nil {
		return nil, ErrBusinessProfileMissing
	}

	job, err := s.jobRepo.GetJobByID(jobID)
	if err != nil {
		return nil, errors.New("Không tìm thấy bài tuyển dụng này")
	}

	if job.BusinessID != business.ID {
		return nil, errors.New("Bạn không có quyền chỉnh sửa bài tuyển dụng của doanh nghiệp khác")
	}

	if input.Title != "" {
		job.Title = input.Title
	}
	if input.Description != "" {
		job.Description = input.Description
	}
	if input.Location != "" {
		job.Location = input.Location
	}
	if input.Salary > 0 {
		job.Salary = input.Salary
	}
	if input.Slots > 0 {
		job.Slots = input.Slots
	}
	if input.WorkingDate != "" {
		job.WorkingDate = input.WorkingDate
	}
	if input.Category != "" {
		job.Category = input.Category
	}
	if input.JobType != "" {
		job.JobType = input.JobType
	}

	if err := s.jobRepo.UpdateJob(job); err != nil {
		return nil, err
	}

	return job, nil
}

func (s *jobService) ToggleJobStatus(userID uint, jobID uint) (*models.Job, error) {
	business, _, err := s.jobRepo.GetBusinessAndUserByUserID(userID)
	if err != nil {
		return nil, ErrBusinessProfileMissing
	}

	job, err := s.jobRepo.GetJobByID(jobID)
	if err != nil {
		return nil, errors.New("Không tìm thấy bài tuyển dụng này")
	}

	if job.BusinessID != business.ID {
		return nil, errors.New("Bạn không có quyền thay đổi bài tuyển dụng của doanh nghiệp khác")
	}

	if job.Status == "closed" {
		job.Status = "approved"
	} else {
		job.Status = "closed"
	}

	if err := s.jobRepo.UpdateJob(job); err != nil {
		return nil, err
	}

	return job, nil
}

func (s *jobService) DeleteJob(userID uint, jobID uint) error {
	business, _, err := s.jobRepo.GetBusinessAndUserByUserID(userID)
	if err != nil {
		return ErrBusinessProfileMissing
	}

	job, err := s.jobRepo.GetJobByID(jobID)
	if err != nil {
		return errors.New("Không tìm thấy bài tuyển dụng này")
	}

	if job.BusinessID != business.ID {
		return errors.New("Bạn không có quyền xóa bài tuyển dụng của doanh nghiệp khác")
	}

	return s.jobRepo.DeleteJob(jobID)
}
