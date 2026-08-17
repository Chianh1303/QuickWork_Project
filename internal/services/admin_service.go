package services

import (
	"errors"
	"math"
	"strings"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"gorm.io/gorm"
)

var (
	ErrPageInvalid           = errors.New("Trang phải lớn hơn hoặc bằng 1")
	ErrLimitInvalid          = errors.New("Số lượng giới hạn phải từ 1 đến 100")
	ErrBusinessIDInvalid     = errors.New("Mã doanh nghiệp không hợp lệ")
	ErrAdminBusinessNotFound = errors.New("Không tìm thấy doanh nghiệp")
	ErrDecisionInvalid       = errors.New("Quyết định duyệt chỉ được là approved hoặc rejected")
	ErrRejectReasonTooShort  = errors.New("Lý do từ chối phải có ít nhất 10 ký tự")
	ErrProfileAlreadyHandled = errors.New("Hồ sơ doanh nghiệp đã được xử lý")
)

type AdminService interface {
	GetDashboardStats() (*dto.AdminDashboardStats, error)
	GetPendingBusinesses(page, limit int, search string) (*dto.PendingBusinessListResponse, error)
	GetBusinessKYBDetail(businessID int) (*dto.BusinessKYBDetail, error)
	ReviewBusinessKYB(businessID int, adminID uint, req dto.ReviewBusinessRequest) error
	GetStudents(page, limit int, search, status string) (*dto.AdminStudentListResponse, error)
	GetStudentDetail(studentID int) (*dto.AdminStudentItem, error)
	UpdateStudentStatus(studentID int, status string) error
	GetBusinesses(page, limit int, search, status string) (*dto.AdminBusinessListResponse, error)
	UpdateBusinessStatus(businessID int, status string) error
	GetTickets(page, limit int, status string) (*dto.AdminTicketListResponse, error)
	GetTicketDetail(ticketID int) (*dto.AdminTicketItem, error)
	ResolveTicket(ticketID int, adminID uint, req dto.ResolveTicketRequest) error
	GetCategories() ([]models.Category, error)
	CreateCategory(req dto.CreateCategoryRequest) (*models.Category, error)
	DeleteCategory(id uint) error
	GetSkills() ([]models.Skill, error)
	CreateSkill(req dto.CreateSkillRequest) (*models.Skill, error)
	DeleteSkill(id uint) error
	GetPendingJobs(page, limit int, search string) (*dto.PendingJobListResponse, error)
	UpdateJobStatus(jobID uint, status string) error
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

func (s *adminService) GetStudents(page, limit int, search, status string) (*dto.AdminStudentListResponse, error) {
	if page < 1 {
		return nil, ErrPageInvalid
	}
	if limit < 1 || limit > 100 {
		return nil, ErrLimitInvalid
	}

	searchClean := strings.TrimSpace(search)
	statusClean := strings.TrimSpace(status)
	items, total, err := s.adminRepo.GetStudents(page, limit, searchClean, statusClean)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	return &dto.AdminStudentListResponse{
		Items: items,
		Pagination: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *adminService) GetStudentDetail(studentID int) (*dto.AdminStudentItem, error) {
	if studentID < 1 {
		return nil, errors.New("student id is invalid")
	}
	detail, err := s.adminRepo.GetStudentDetail(studentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStudentNotFound
		}
		return nil, err
	}
	return detail, nil
}

func (s *adminService) UpdateStudentStatus(studentID int, status string) error {
	if studentID < 1 {
		return errors.New("student id is invalid")
	}
	statusClean := strings.TrimSpace(status)
	if statusClean != "approved" && statusClean != "locked" {
		return errors.New("status chỉ được là approved hoặc locked")
	}
	err := s.adminRepo.UpdateStudentStatus(studentID, statusClean)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrStudentNotFound
		}
		return err
	}
	return nil
}

func (s *adminService) GetBusinesses(page, limit int, search, status string) (*dto.AdminBusinessListResponse, error) {
	if page < 1 {
		return nil, ErrPageInvalid
	}
	if limit < 1 || limit > 100 {
		return nil, ErrLimitInvalid
	}

	searchClean := strings.TrimSpace(search)
	statusClean := strings.TrimSpace(status)
	items, total, err := s.adminRepo.GetBusinesses(page, limit, searchClean, statusClean)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	return &dto.AdminBusinessListResponse{
		Items: items,
		Pagination: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *adminService) UpdateBusinessStatus(businessID int, status string) error {
	if businessID < 1 {
		return ErrBusinessIDInvalid
	}
	statusClean := strings.TrimSpace(status)
	if statusClean != "approved" && statusClean != "locked" {
		return errors.New("status chỉ được là approved hoặc locked")
	}
	err := s.adminRepo.UpdateBusinessStatus(businessID, statusClean)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAdminBusinessNotFound
		}
		return err
	}
	return nil
}

func (s *adminService) GetTickets(page, limit int, status string) (*dto.AdminTicketListResponse, error) {
	if page < 1 {
		return nil, ErrPageInvalid
	}
	if limit < 1 || limit > 100 {
		return nil, ErrLimitInvalid
	}

	statusClean := strings.TrimSpace(status)
	items, total, err := s.adminRepo.GetTickets(page, limit, statusClean)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	return &dto.AdminTicketListResponse{
		Items: items,
		Pagination: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *adminService) GetTicketDetail(ticketID int) (*dto.AdminTicketItem, error) {
	if ticketID < 1 {
		return nil, errors.New("ticket id is invalid")
	}
	detail, err := s.adminRepo.GetTicketDetail(ticketID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Không tìm thấy khiếu nại")
		}
		return nil, err
	}
	return detail, nil
}

func (s *adminService) ResolveTicket(ticketID int, adminID uint, req dto.ResolveTicketRequest) error {
	if ticketID < 1 {
		return errors.New("ticket id is invalid")
	}
	verdict := strings.TrimSpace(req.Verdict)
	if len([]rune(verdict)) < 5 {
		return errors.New("Phán quyết của Admin phải có ít nhất 5 ký tự")
	}
	status := strings.TrimSpace(req.Status)
	if status != "resolved" && status != "rejected" {
		status = "resolved"
	}
	return s.adminRepo.ResolveTicket(ticketID, adminID, verdict, status)
}

func (s *adminService) GetCategories() ([]models.Category, error) {
	return s.adminRepo.GetCategories()
}

func (s *adminService) CreateCategory(req dto.CreateCategoryRequest) (*models.Category, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("Tên ngành nghề không được để trống")
	}
	cat := &models.Category{
		Name:        name,
		Description: strings.TrimSpace(req.Description),
	}
	if err := s.adminRepo.CreateCategory(cat); err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *adminService) DeleteCategory(id uint) error {
	if id < 1 {
		return errors.New("ID ngành nghề không hợp lệ")
	}
	return s.adminRepo.DeleteCategory(id)
}

func (s *adminService) GetSkills() ([]models.Skill, error) {
	return s.adminRepo.GetSkills()
}

func (s *adminService) CreateSkill(req dto.CreateSkillRequest) (*models.Skill, error) {
	name := strings.TrimSpace(req.Name)
	if name == "" {
		return nil, errors.New("Tên kỹ năng không được để trống")
	}
	category := strings.TrimSpace(req.Category)
	if category == "" {
		category = "General"
	}
	skill := &models.Skill{
		Name:     name,
		Category: category,
	}
	if err := s.adminRepo.CreateSkill(skill); err != nil {
		return nil, err
	}
	return skill, nil
}

func (s *adminService) DeleteSkill(id uint) error {
	if id < 1 {
		return errors.New("ID kỹ năng không hợp lệ")
	}
	return s.adminRepo.DeleteSkill(id)
}

func (s *adminService) GetPendingJobs(page, limit int, search string) (*dto.PendingJobListResponse, error) {
	if page < 1 {
		return nil, ErrPageInvalid
	}
	if limit < 1 || limit > 100 {
		return nil, ErrLimitInvalid
	}

	searchClean := strings.TrimSpace(search)
	items, total, err := s.adminRepo.GetPendingJobs(page, limit, searchClean)
	if err != nil {
		return nil, err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	return &dto.PendingJobListResponse{
		Items: items,
		Pagination: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *adminService) UpdateJobStatus(jobID uint, status string) error {
	if jobID < 1 {
		return errors.New("job id is invalid")
	}
	statusClean := strings.TrimSpace(status)
	if statusClean != "approved" && statusClean != "rejected" && statusClean != "closed" {
		return errors.New("status chỉ được là approved, rejected hoặc closed")
	}
	return s.adminRepo.UpdateJobStatus(jobID, statusClean)
}

