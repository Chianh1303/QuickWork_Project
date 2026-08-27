package repositories

import (
	"QuickWork/internal/models"
	"errors"

	"gorm.io/gorm"
)

type SavedJobRepository interface {
	SaveJob(studentID uint, jobID uint) error
	UnsaveJob(studentID uint, jobID uint) error
	GetSavedJobsByStudentID(studentID uint) ([]models.SavedJob, error)
	IsJobSaved(studentID uint, jobID uint) (bool, error)
	GetStudentByUserID(userID uint) (*models.Student, error)
}

type savedJobRepository struct {
	db *gorm.DB
}

func NewSavedJobRepository(db *gorm.DB) SavedJobRepository {
	return &savedJobRepository{db: db}
}

func (r *savedJobRepository) SaveJob(studentID uint, jobID uint) error {
	var existing models.SavedJob
	err := r.db.Where("student_id = ? AND job_id = ?", studentID, jobID).First(&existing).Error
	if err == nil {
		return nil // Already saved
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	saved := models.SavedJob{
		StudentID: studentID,
		JobID:     jobID,
	}
	return r.db.Create(&saved).Error
}

func (r *savedJobRepository) UnsaveJob(studentID uint, jobID uint) error {
	return r.db.Where("student_id = ? AND job_id = ?", studentID, jobID).Delete(&models.SavedJob{}).Error
}

func (r *savedJobRepository) GetSavedJobsByStudentID(studentID uint) ([]models.SavedJob, error) {
	var list []models.SavedJob
	err := r.db.Preload("Job").Preload("Job.Business").Where("student_id = ?", studentID).Order("created_at desc").Find(&list).Error
	return list, err
}

func (r *savedJobRepository) IsJobSaved(studentID uint, jobID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.SavedJob{}).Where("student_id = ? AND job_id = ?", studentID, jobID).Count(&count).Error
	return count > 0, err
}

func (r *savedJobRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	err := r.db.Where("user_id = ?", userID).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
