package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewsByRevieweeID(userID uint) ([]models.Review, error)
	GetReviewsByApplicationID(appID uint) ([]models.Review, error)
	GetApplicationByID(appID uint) (*models.Application, error)
	GetExistingReview(appID, reviewerID uint) (*models.Review, error)
	CreateReview(review *models.Review) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetReviewsByRevieweeID(userID uint) ([]models.Review, error) {
	var reviews []models.Review
	if err := r.db.
		Where("reviewee_id = ?", userID).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) GetReviewsByApplicationID(appID uint) ([]models.Review, error) {
	var reviews []models.Review
	if err := r.db.
		Where("application_id = ?", appID).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) GetApplicationByID(appID uint) (*models.Application, error) {
	var app models.Application
	if err := r.db.
		Preload("Job.Business").
		Preload("Student").
		First(&app, appID).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *reviewRepository) GetExistingReview(appID, reviewerID uint) (*models.Review, error) {
	var existing models.Review
	if err := r.db.Where("application_id = ? AND reviewer_id = ?", appID, reviewerID).First(&existing).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}
