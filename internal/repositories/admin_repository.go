package repositories

import (
	"fmt"
	"time"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminRepository interface {
	GetDashboardStats() (*dto.AdminDashboardStats, error)
	GetPendingBusinesses(page, limit int, search string) ([]dto.PendingBusinessItem, int64, error)
	GetBusinessKYBDetail(businessID int) (*dto.BusinessKYBDetail, error)
	ReviewBusinessKYB(businessID int, adminID uint, decision, rejectReason string) error
	GetStudents(page, limit int, search, status string) ([]dto.AdminStudentItem, int64, error)
	GetStudentDetail(studentID int) (*dto.AdminStudentItem, error)
	UpdateStudentStatus(studentID int, status string) error
	GetBusinesses(page, limit int, search, status string) ([]dto.AdminBusinessItem, int64, error)
	UpdateBusinessStatus(businessID int, status string) error
	GetTickets(page, limit int, status string) ([]dto.AdminTicketItem, int64, error)
	GetTicketDetail(ticketID int) (*dto.AdminTicketItem, error)
	ResolveTicket(ticketID int, adminID uint, verdict, status string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetDashboardStats() (*dto.AdminDashboardStats, error) {
	var stats dto.AdminDashboardStats

	if err := r.db.Model(&models.User{}).Where("role = ?", "student").Count(&stats.TotalStudents).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&models.User{}).Where("role = ?", "business").Count(&stats.TotalBusinesses).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&models.User{}).Where("role = ? AND status = ?", "business", "pending").Count(&stats.PendingBusinesses).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&models.Job{}).Count(&stats.TotalJobs).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&models.Job{}).Where("status = ?", "pending").Count(&stats.PendingJobs).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&models.WalletTransaction{}).
		Where("type = ?", "salary").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&stats.TotalDisbursed).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *adminRepository) GetPendingBusinesses(page, limit int, search string) ([]dto.PendingBusinessItem, int64, error) {
	offset := (page - 1) * limit
	buildQuery := func() *gorm.DB {
		query := r.db.
			Table("businesses").
			Joins("JOIN users ON users.id = businesses.user_id").
			Where("users.role = ? AND users.status = ?", "business", "pending")

		if search != "" {
			keyword := "%" + search + "%"
			query = query.Where(
				"(businesses.company_name LIKE ? OR businesses.tax_code LIKE ? OR users.email LIKE ?)",
				keyword,
				keyword,
				keyword,
			)
		}

		return query
	}

	var total int64
	if err := buildQuery().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	items := make([]dto.PendingBusinessItem, 0)
	if err := buildQuery().
		Select(`
			businesses.id AS business_id,
			businesses.company_name,
			businesses.tax_code,
			users.email,
			users.status,
			businesses.created_at
		`).
		Order("businesses.created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *adminRepository) GetBusinessKYBDetail(businessID int) (*dto.BusinessKYBDetail, error) {
	var detail dto.BusinessKYBDetail
	err := r.db.
		Table("businesses").
		Joins("JOIN users ON users.id = businesses.user_id").
		Where("businesses.id = ? AND users.role = ?", businessID, "business").
		Select(`
			businesses.id AS business_id,
			businesses.user_id,
			businesses.company_name,
			businesses.tax_code,
			users.email,
			businesses.phone,
			businesses.address,
			businesses.logo_url,
			users.status,
			businesses.is_verified,
			businesses.reject_reason,
			businesses.created_at,
			businesses.reviewed_at
		`).
		Scan(&detail).Error
	if err != nil {
		return nil, err
	}
	if detail.BusinessID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &detail, nil
}

func (r *adminRepository) ReviewBusinessKYB(businessID int, adminID uint, decision, rejectReason string) error {
	now := time.Now()
	return r.db.Transaction(func(tx *gorm.DB) error {
		var business models.Business
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&business, uint(businessID)).Error; err != nil {
			return err
		}

		var user models.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ? AND role = ?", business.UserID, "business").
			First(&user).Error; err != nil {
			return err
		}

		if user.Status != "pending" {
			return fmt.Errorf("PROFILE_ALREADY_PROCESSED")
		}

		nextStatus := "approved"
		isVerified := true
		reason := ""
		if decision == "approved" {
			nextStatus = "approved"
			isVerified = true
			reason = ""
		} else {
			nextStatus = "rejected"
			isVerified = false
			reason = rejectReason
		}

		if err := tx.Model(&user).Update("status", nextStatus).Error; err != nil {
			return err
		}

		if err := tx.Model(&business).Updates(map[string]interface{}{
			"is_verified":   isVerified,
			"reject_reason": reason,
			"reviewed_by":   adminID,
			"reviewed_at":   now,
		}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *adminRepository) GetStudents(page, limit int, search, status string) ([]dto.AdminStudentItem, int64, error) {
	offset := (page - 1) * limit
	buildQuery := func() *gorm.DB {
		query := r.db.
			Table("students").
			Joins("JOIN users ON users.id = students.user_id").
			Where("users.role = ?", "student")

		if status != "" {
			query = query.Where("users.status = ?", status)
		}

		if search != "" {
			keyword := "%" + search + "%"
			query = query.Where(
				"(students.full_name LIKE ? OR students.phone LIKE ? OR users.email LIKE ?)",
				keyword,
				keyword,
				keyword,
			)
		}

		return query
	}

	var total int64
	if err := buildQuery().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	items := make([]dto.AdminStudentItem, 0)
	if err := buildQuery().
		Select(`
			students.id AS student_id,
			students.user_id,
			students.full_name,
			users.email,
			students.phone,
			students.gender,
			students.avatar_url,
			students.skills,
			students.cv_url,
			users.status,
			users.created_at
		`).
		Order("students.id DESC").
		Limit(limit).
		Offset(offset).
		Scan(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *adminRepository) GetStudentDetail(studentID int) (*dto.AdminStudentItem, error) {
	var detail dto.AdminStudentItem
	err := r.db.
		Table("students").
		Joins("JOIN users ON users.id = students.user_id").
		Where("students.id = ? AND users.role = ?", studentID, "student").
		Select(`
			students.id AS student_id,
			students.user_id,
			students.full_name,
			users.email,
			students.phone,
			students.gender,
			students.avatar_url,
			students.skills,
			students.cv_url,
			users.status,
			users.created_at
		`).
		Scan(&detail).Error
	if err != nil {
		return nil, err
	}
	if detail.StudentID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &detail, nil
}

func (r *adminRepository) UpdateStudentStatus(studentID int, status string) error {
	var student models.Student
	if err := r.db.First(&student, studentID).Error; err != nil {
		return err
	}
	return r.db.Model(&models.User{}).Where("id = ? AND role = ?", student.UserID, "student").Update("status", status).Error
}

func (r *adminRepository) GetBusinesses(page, limit int, search, status string) ([]dto.AdminBusinessItem, int64, error) {
	offset := (page - 1) * limit
	buildQuery := func() *gorm.DB {
		query := r.db.
			Table("businesses").
			Joins("JOIN users ON users.id = businesses.user_id").
			Where("users.role = ?", "business")

		if status != "" {
			query = query.Where("users.status = ?", status)
		}

		if search != "" {
			keyword := "%" + search + "%"
			query = query.Where(
				"(businesses.company_name LIKE ? OR businesses.tax_code LIKE ? OR users.email LIKE ? OR businesses.phone LIKE ?)",
				keyword,
				keyword,
				keyword,
				keyword,
			)
		}

		return query
	}

	var total int64
	if err := buildQuery().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	items := make([]dto.AdminBusinessItem, 0)
	if err := buildQuery().
		Select(`
			businesses.id AS business_id,
			businesses.user_id,
			businesses.company_name,
			businesses.tax_code,
			users.email,
			businesses.phone,
			businesses.address,
			businesses.logo_url,
			users.status,
			businesses.is_verified,
			businesses.created_at,
			(SELECT COUNT(*) FROM jobs WHERE jobs.business_id = businesses.id) AS job_count
		`).
		Order("businesses.id DESC").
		Limit(limit).
		Offset(offset).
		Scan(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *adminRepository) UpdateBusinessStatus(businessID int, status string) error {
	var business models.Business
	if err := r.db.First(&business, businessID).Error; err != nil {
		return err
	}
	return r.db.Model(&models.User{}).Where("id = ? AND role = ?", business.UserID, "business").Update("status", status).Error
}

func (r *adminRepository) GetTickets(page, limit int, status string) ([]dto.AdminTicketItem, int64, error) {
	offset := (page - 1) * limit
	buildQuery := func() *gorm.DB {
		query := r.db.Table("tickets")
		if status != "" {
			query = query.Where("status = ?", status)
		}
		return query
	}

	var total int64
	if err := buildQuery().Count(&total).Error; err != nil {
		return nil, 0, err
	}

	items := make([]dto.AdminTicketItem, 0)
	if err := buildQuery().
		Select(`
			tickets.id AS ticket_id,
			tickets.application_id,
			tickets.reporter_id,
			u1.email AS reporter_email,
			u1.role AS reporter_role,
			tickets.target_id,
			u2.email AS target_email,
			tickets.reason,
			tickets.description,
			tickets.status,
			tickets.verdict,
			tickets.resolved_at,
			tickets.created_at
		`).
		Joins("LEFT JOIN users u1 ON u1.id = tickets.reporter_id").
		Joins("LEFT JOIN users u2 ON u2.id = tickets.target_id").
		Order("tickets.id DESC").
		Limit(limit).
		Offset(offset).
		Scan(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *adminRepository) GetTicketDetail(ticketID int) (*dto.AdminTicketItem, error) {
	var detail dto.AdminTicketItem
	err := r.db.Table("tickets").
		Select(`
			tickets.id AS ticket_id,
			tickets.application_id,
			tickets.reporter_id,
			u1.email AS reporter_email,
			u1.role AS reporter_role,
			tickets.target_id,
			u2.email AS target_email,
			tickets.reason,
			tickets.description,
			tickets.status,
			tickets.verdict,
			tickets.resolved_at,
			tickets.created_at
		`).
		Joins("LEFT JOIN users u1 ON u1.id = tickets.reporter_id").
		Joins("LEFT JOIN users u2 ON u2.id = tickets.target_id").
		Where("tickets.id = ?", ticketID).
		Scan(&detail).Error
	if err != nil {
		return nil, err
	}
	if detail.TicketID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &detail, nil
}

func (r *adminRepository) ResolveTicket(ticketID int, adminID uint, verdict, status string) error {
	now := time.Now()
	return r.db.Model(&models.Ticket{}).Where("id = ?", ticketID).Updates(map[string]interface{}{
		"status":      status,
		"verdict":     verdict,
		"resolved_by": adminID,
		"resolved_at": now,
	}).Error
}

