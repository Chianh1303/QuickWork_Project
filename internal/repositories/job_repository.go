package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type JobRepository interface {
	GetAvailableJobs(search, location, category, jobType, maxSalary string) ([]models.Job, error)
	GetBusinessAndUserByUserID(userID uint) (*models.Business, *models.User, error)
	CreateJob(job *models.Job) error
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
		Where("status = ?", "approved").
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
