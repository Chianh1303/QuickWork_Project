package dto

import "time"

// PendingBusinessItem chỉ chứa dữ liệu cần thiết cho màn hình danh sách.
// Không trả trực tiếp toàn bộ Business entity.
type PendingBusinessItem struct {
	BusinessID  uint      `json:"business_id"`
	CompanyName string    `json:"company_name"`
	TaxCode     string    `json:"tax_code"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type PendingBusinessListResponse struct {
	Items      []PendingBusinessItem `json:"items"`
	Pagination PaginationMeta        `json:"pagination"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type AdminDashboardStats struct {
	TotalStudents     int64   `json:"total_students"`
	TotalBusinesses   int64   `json:"total_businesses"`
	PendingBusinesses int64   `json:"pending_businesses"`
	TotalJobs         int64   `json:"total_jobs"`
	PendingJobs       int64   `json:"pending_jobs"`
	TotalDisbursed    float64 `json:"total_disbursed"`
}

type BusinessKYBDetail struct {
	BusinessID   uint       `json:"business_id"`
	UserID       uint       `json:"user_id"`
	CompanyName  string     `json:"company_name"`
	TaxCode      string     `json:"tax_code"`
	Email        string     `json:"email"`
	Phone        string     `json:"phone"`
	Address      string     `json:"address"`
	LogoURL      string     `json:"logo_url"`
	Status       string     `json:"status"`
	IsVerified   bool       `json:"is_verified"`
	RejectReason string     `json:"reject_reason"`
	CreatedAt    time.Time  `json:"created_at"`
	ReviewedAt   *time.Time `json:"reviewed_at"`
}

type ReviewBusinessRequest struct {
	Decision     string `json:"decision"`
	RejectReason string `json:"reject_reason"`
}

type AdminBusinessItem struct {
	BusinessID  uint      `json:"business_id"`
	UserID      uint      `json:"user_id"`
	CompanyName string    `json:"company_name"`
	TaxCode     string    `json:"tax_code"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	LogoURL     string    `json:"logo_url"`
	Status      string    `json:"status"` // approved, pending, locked, rejected
	IsVerified  bool      `json:"is_verified"`
	JobCount    int64     `json:"job_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type AdminBusinessListResponse struct {
	Items      []AdminBusinessItem `json:"items"`
	Pagination PaginationMeta      `json:"pagination"`
}

type UpdateBusinessStatusRequest struct {
	Status string `json:"status"` // approved, locked
}
