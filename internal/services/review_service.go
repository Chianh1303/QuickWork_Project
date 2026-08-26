package services

import (
	"errors"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrMissingApplicationID = errors.New("Thiếu application_id")
	ErrRatingOutOfBounds     = errors.New("Rating phải từ 1 đến 5")
	ErrReviewAppNotFound     = errors.New("Không tìm thấy Application")
	ErrJobNotFinished        = errors.New("Công việc chưa hoàn thành.")
	ErrNoPermissionToReview  = errors.New("Bạn không có quyền đánh giá.")
	ErrAlreadyReviewed       = errors.New("Bạn đã đánh giá rồi.")
)

type ReviewService interface {
	GetReviewsByUser(userID uint) ([]models.Review, float64, int, error)
	GetReviewsByApplication(appID uint) ([]models.Review, error)
	CreateReview(userID uint, input dto.CreateReviewInput) (*models.Review, error)
}

type reviewService struct {
	reviewRepo repositories.ReviewRepository
}

func NewReviewService(reviewRepo repositories.ReviewRepository) ReviewService {
	return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) GetReviewsByUser(userID uint) ([]models.Review, float64, int, error) {
	reviews, err := s.reviewRepo.GetReviewsByRevieweeID(userID)
	if err != nil {
		return nil, 0, 0, err
	}

	total := len(reviews)
	var average float64
	if total > 0 {
		var sum int
		for _, r := range reviews {
			sum += r.Rating
		}
		average = float64(sum) / float64(total)
	}

	return reviews, average, total, nil
}

func (s *reviewService) GetReviewsByApplication(appID uint) ([]models.Review, error) {
	return s.reviewRepo.GetReviewsByApplicationID(appID)
}

func (s *reviewService) CreateReview(userID uint, input dto.CreateReviewInput) (*models.Review, error) {
	if input.ApplicationID == 0 {
		return nil, ErrMissingApplicationID
	}
	if input.Rating < 1 || input.Rating > 5 {
		return nil, ErrRatingOutOfBounds
	}

	app, err := s.reviewRepo.GetApplicationByID(input.ApplicationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrReviewAppNotFound
		}
		return nil, err
	}

	if app.Status != "paid" {
		return nil, ErrJobNotFinished
	}

	var reviewerID uint
	var revieweeID uint
	var reviewType string

	if app.Student.UserID == userID {
		reviewerID = app.Student.UserID
		revieweeID = app.Job.Business.UserID
		reviewType = "business"
	}
	if app.Job.Business.UserID == userID {
		reviewerID = app.Job.Business.UserID
		revieweeID = app.Student.UserID
		reviewType = "student"
	}
	if reviewerID == 0 {
		return nil, ErrNoPermissionToReview
	}

	existing, err := s.reviewRepo.GetExistingReview(input.ApplicationID, reviewerID)
	if err == nil && existing != nil {
		return nil, ErrAlreadyReviewed
	}

	review := &models.Review{
		ApplicationID: input.ApplicationID,
		ReviewerID:    reviewerID,
		RevieweeID:    revieweeID,
		ReviewType:    reviewType,
		Rating:        input.Rating,
		Comment:       input.Comment,
	}

	if err := s.reviewRepo.CreateReview(review); err != nil {
		return nil, err
	}

	return review, nil
}
