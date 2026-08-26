package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetStudentByUserID(userID uint) (*models.Student, error)
	SaveStudent(student *models.Student) error
	GetBusinessByUserID(userID uint) (*models.Business, error)
	SaveBusiness(business *models.Business) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *profileRepository) SaveStudent(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *profileRepository) GetBusinessByUserID(userID uint) (*models.Business, error) {
	var business models.Business
	if err := r.db.Where("user_id = ?", userID).First(&business).Error; err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *profileRepository) SaveBusiness(business *models.Business) error {
	return r.db.Save(business).Error
}
