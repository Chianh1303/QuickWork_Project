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
