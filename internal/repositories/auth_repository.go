package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUserWithProfile(user *models.User, student *models.Student, business *models.Business) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) CreateUserWithProfile(user *models.User, student *models.Student, business *models.Business) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		if student != nil {
			student.UserID = user.ID
			if err := tx.Create(student).Error; err != nil {
				return err
			}
		}
		if business != nil {
			business.UserID = user.ID
			if err := tx.Create(business).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
