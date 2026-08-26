package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetAvailableJobs(search, location, category, jobType, maxSalary string) ([]models.Job, error)
	GetBusinessAndUserByUserID(userID uint) (*models.Business, *models.User, error)
	CreateJob(job *models.Job) error
	GetJobsByBusinessID(businessID uint) ([]models.Job, error)
	GetJobByID(jobID uint) (*models.Job, error)
	UpdateJob(job *models.Job) error
	DeleteJob(jobID uint) error
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) GetAvailableJobs(search, location, category, jobType, maxSalary string) ([]models.Job, error) {
	var jobs []models.Job
	query := r.db.Model(&models.Job{})

	if search != "" {
		query = query.Where("LOWER(title) LIKE LOWER(?)", "%"+search+"%")
	}
	if location != "" && location != "all" {
		query = query.Where("LOWER(location) LIKE LOWER(?)", "%"+location+"%")
	}
	if category != "" && category != "all" {
		query = query.Where("LOWER(category) = LOWER(?)", category)
	}
	if jobType != "" && jobType != "all" {
		query = query.Where("LOWER(job_type) = LOWER(?)", jobType)
	}
	if maxSalary != "" {
		query = query.Where("salary >= ?", maxSalary)
	}

	if err := query.
		Preload("Business").
		Where("status IN ?", []string{"approved", "open"}).
		Order("id DESC").
		Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (r *jobRepository) GetBusinessAndUserByUserID(userID uint) (*models.Business, *models.User, error) {
	var business models.Business
	if err := r.db.Where("user_id = ?", userID).First(&business).Error; err != nil {
		return nil, nil, err
	}
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return &business, nil, err
	}
	return &business, &user, nil
}

func (r *jobRepository) CreateJob(job *models.Job) error {
	return r.db.Create(job).Error
}

func (r *jobRepository) GetJobsByBusinessID(businessID uint) ([]models.Job, error) {
	var jobs []models.Job
	err := r.db.Where("business_id = ?", businessID).Order("id DESC").Find(&jobs).Error
	return jobs, err
}

func (r *jobRepository) GetJobByID(jobID uint) (*models.Job, error) {
	var job models.Job
	err := r.db.First(&job, jobID).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *jobRepository) UpdateJob(job *models.Job) error {
	return r.db.Save(job).Error
}

func (r *jobRepository) DeleteJob(jobID uint) error {
	return r.db.Delete(&models.Job{}, jobID).Error
}
