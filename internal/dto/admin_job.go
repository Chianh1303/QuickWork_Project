package dto

import "time"

type PendingJobDTO struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	CompanyName string    `json:"company_name"`
	LogoURL     string    `json:"logo_url"`
	TaxCode     string    `json:"tax_code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Salary      float64   `json:"salary"`
	Slots       int       `json:"slots"`
	WorkingDate string    `json:"working_date"`
	Category    string    `json:"category"`
	JobType     string    `json:"job_type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type PendingJobListResponse struct {
	Items      []PendingJobDTO `json:"items"`
	Pagination PaginationMeta  `json:"pagination"`
}

type UpdateJobStatusRequest struct {
	Status string `json:"status" validate:"required"` // approved, rejected, closed
}
