package services

import (
	"errors"
	"math"
	"strings"

	"QuickWork/internal/dto"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrPageInvalid           = errors.New("page must be greater than or equal to 1")
	ErrLimitInvalid          = errors.New("limit must be between 1 and 100")
	ErrBusinessIDInvalid     = errors.New("business id is invalid")
	ErrAdminBusinessNotFound = errors.New("business not found")
	ErrDecisionInvalid       = errors.New("decision chỉ được là approved hoặc rejected")
	ErrRejectReasonTooShort  = errors.New("Lý do từ chối phải có ít nhất 10 ký tự")
	ErrProfileAlreadyHandled = errors.New("Hồ sơ doanh nghiệp đã được xử lý")
)

type AdminService interface {
	GetDashboardStats() (*dto.AdminDashboardStats, error)
	GetPendingBusinesses(page, limit int, search string) (*dto.PendingBusinessListResponse, error)
	GetBusinessKYBDetail(businessID int) (*dto.BusinessKYBDetail, error)
	ReviewBusinessKYB(businessID int, adminID uint, req dto.ReviewBusinessRequest) error
}

type adminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService(adminRepo repositories.AdminRepository) AdminService {
	return &adminService{adminRepo: adminRepo}
}

func (s *adminService) GetDashboardStats() (*dto.AdminDashboardStats, error) {
	return s.adminRepo.GetDashboardStats()
}

func (s *adminService) GetPendingBusinesses(page, limit int, search string) (*dto.PendingBusinessListResponse, error) {
	if page < 1 {
		return nil, ErrPageInvalid
	}
	if limit < 1 || limit > 100 {
		return nil, ErrLimitInvalid
	}

	searchClean := strings.TrimSpace(search)
	items, total, err := s.adminRepo.GetPendingBusinesses(page, limit, searchClean)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	return &dto.PendingBusinessListResponse{
		Items: items,
		Pagination: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *adminService) GetBusinessKYBDetail(businessID int) (*dto.BusinessKYBDetail, error) {
	if businessID < 1 {
		return nil, ErrBusinessIDInvalid
	}
	detail, err := s.adminRepo.GetBusinessKYBDetail(businessID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAdminBusinessNotFound
		}
		return nil, err
	}
	return detail, nil
}

func (s *adminService) ReviewBusinessKYB(businessID int, adminID uint, req dto.ReviewBusinessRequest) error {
	if businessID < 1 {
		return ErrBusinessIDInvalid
	}

	decision := strings.TrimSpace(req.Decision)
	rejectReason := strings.TrimSpace(req.RejectReason)

	if decision != "approved" && decision != "rejected" {
		return ErrDecisionInvalid
	}
	if decision == "rejected" && len([]rune(rejectReason)) < 10 {
		return ErrRejectReasonTooShort
	}

	err := s.adminRepo.ReviewBusinessKYB(businessID, adminID, decision, rejectReason)
	if err != nil {
		if err.Error() == "PROFILE_ALREADY_PROCESSED" {
			return ErrProfileAlreadyHandled
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAdminBusinessNotFound
		}
		return err
	}

	return nil
}
