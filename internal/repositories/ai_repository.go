package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type AIRepository interface {
	GetStudentByUserID(userID uint) (*models.Student, error)
	GetJobByID(jobID uint) (*models.Job, error)
}

type aiRepository struct {
	db *gorm.DB
}

func NewAIRepository(db *gorm.DB) AIRepository {
	return &aiRepository{db: db}
}

func (r *aiRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *aiRepository) GetJobByID(jobID uint) (*models.Job, error) {
	var job models.Job
	if err := r.db.First(&job, jobID).Error; err != nil {
		return nil, err
	}
	return &job, nil
}
