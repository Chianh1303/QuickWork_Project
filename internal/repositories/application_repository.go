package repositories

import (
	"fmt"
	"time"

	"QuickWork/internal/models"

	"gorm.io/gorm"
)

type ApplicationRepository interface {
	GetStudentByUserID(userID uint) (*models.Student, error)
	GetStudentApplications(studentID uint) ([]models.Application, error)
	GetJobByID(jobID uint) (*models.Job, error)
	GetApplicationByJobAndStudent(jobID, studentID uint) (*models.Application, error)
	CreateApplication(app *models.Application) error
	GetApplicationByIDAndStudent(appID, studentID uint) (*models.Application, error)
	DeleteApplication(app *models.Application) error
	GetBusinessByUserID(userID uint) (*models.Business, error)
	GetEmployerApplications(businessID uint) ([]models.Application, error)
	GetApplicationByID(appID uint) (*models.Application, error)
	GetStudentByID(studentID uint) (*models.Student, error)
	SaveApplication(app *models.Application) error
	BusinessCompleteJobAndPay(appID, businessID uint) (*models.Application, error)
}

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepository{db: db}
}

func (r *applicationRepository) GetStudentByUserID(userID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *applicationRepository) GetStudentApplications(studentID uint) ([]models.Application, error) {
	var apps []models.Application
	if err := r.db.
		Preload("Job").
		Preload("Job.Business").
		Preload("Job.Business.User").
		Where("student_id = ?", studentID).
		Find(&apps).Error; err != nil {
		return nil, err
	}
	return apps, nil
}

func (r *applicationRepository) GetJobByID(jobID uint) (*models.Job, error) {
	var job models.Job
	if err := r.db.Preload("Business").First(&job, jobID).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *applicationRepository) GetApplicationByJobAndStudent(jobID, studentID uint) (*models.Application, error) {
	var app models.Application
	if err := r.db.Where("job_id = ? AND student_id = ?", jobID, studentID).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *applicationRepository) CreateApplication(app *models.Application) error {
	return r.db.Create(app).Error
}

func (r *applicationRepository) GetApplicationByIDAndStudent(appID, studentID uint) (*models.Application, error) {
	var app models.Application
	if err := r.db.Preload("Job").Where("id = ? AND student_id = ?", appID, studentID).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *applicationRepository) DeleteApplication(app *models.Application) error {
	return r.db.Delete(app).Error
}

func (r *applicationRepository) GetBusinessByUserID(userID uint) (*models.Business, error) {
	var business models.Business
	if err := r.db.Where("user_id = ?", userID).First(&business).Error; err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *applicationRepository) GetEmployerApplications(businessID uint) ([]models.Application, error) {
	var apps []models.Application
	err := r.db.
		Preload("Student").
		Preload("Student.User").
		Preload("Job").
		Preload("Job.Business").
		Preload("Job.Business.User").
		Joins("JOIN jobs ON jobs.id = applications.job_id").
		Where("jobs.business_id = ?", businessID).
		Find(&apps).Error
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (r *applicationRepository) GetApplicationByID(appID uint) (*models.Application, error) {
	var app models.Application
	if err := r.db.Preload("Job").Preload("Student").First(&app, appID).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *applicationRepository) GetStudentByID(studentID uint) (*models.Student, error) {
	var student models.Student
	if err := r.db.First(&student, studentID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *applicationRepository) SaveApplication(app *models.Application) error {
	return r.db.Save(app).Error
}

func (r *applicationRepository) BusinessCompleteJobAndPay(appID, businessID uint) (*models.Application, error) {
	var app models.Application
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Job").Preload("Student").First(&app, appID).Error; err != nil {
			return err
		}
		if app.Job.BusinessID != businessID {
			return fmt.Errorf("FORBIDDEN_BUSINESS_OWNERSHIP")
		}
		if app.Status != "student_completed" {
			return fmt.Errorf("STUDENT_NOT_COMPLETED")
		}

		app.Status = "paid"
		if err := tx.Save(&app).Error; err != nil {
			return err
		}

		var studentUser models.User
		if err := tx.First(&studentUser, app.Student.UserID).Error; err != nil {
			return err
		}

		var wallet models.Wallet
		if err := tx.Where("user_id = ?", studentUser.ID).First(&wallet).Error; err != nil {
			wallet = models.Wallet{
				UserID:  studentUser.ID,
				Balance: 0,
			}
			if err := tx.Create(&wallet).Error; err != nil {
				return err
			}
		}

		salaryAmount := app.Job.Salary
		wallet.Balance += salaryAmount
		if err := tx.Save(&wallet).Error; err != nil {
			return err
		}

		studentUser.Balance += salaryAmount
		if err := tx.Save(&studentUser).Error; err != nil {
			return err
		}

		transaction := models.WalletTransaction{
			WalletID:    wallet.ID,
			Amount:      salaryAmount,
			Type:        "salary",
			Description: fmt.Sprintf("Nhận lương cho công việc: %s", app.Job.Title),
			CreatedAt:   time.Now(),
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &app, nil
}
