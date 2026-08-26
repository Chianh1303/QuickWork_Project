package models

import "time"

type Application struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	JobID             uint       `gorm:"not null" json:"job_id"`
	StudentID         uint       `gorm:"not null" json:"student_id"`
	Status            string     `gorm:"type:varchar(20);default:'pending'" json:"status"` // Sửa thành pending cho khớp Nuxt
	CoverNote         string     `gorm:"type:text" json:"cover_note"`                      // 🌟 Thêm trường Cover Note
	Job               Job        `gorm:"foreignKey:JobID" json:"job,omitempty"`
	Student           Student    `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	OfferSalary       string     `json:"offer_salary"`     // Mức lương offer cho sinh viên
	OfferStartDate    string     `json:"offer_start_date"` // Ngày bắt đầu làm việc
	OfferMessage      string     `json:"offer_message"`
	StudentCompleted   bool       `gorm:"default:false" json:"student_completed"`
	BusinessCompleted  bool       `gorm:"default:false" json:"business_completed"`
	CompletionNote     string     `gorm:"type:text" json:"completion_note"`
	CompletionProofUrl string     `gorm:"type:varchar(500)" json:"completion_proof_url"`
	SubmittedAt        *time.Time `json:"submitted_at"`
	CompletedAt        *time.Time `json:"completed_at"`
	PaidAt             *time.Time `json:"paid_at"`
	PaymentStatus      string     `gorm:"type:varchar(20);default:'unpaid'" json:"payment_status"`
}

// Struct bóc tách dữ liệu body gửi lên
type ApplyJobInput struct {
	JobID     uint   `json:"job_id"`
	CoverNote string `json:"cover_note"` // 🌟 Đón đầu Cover Note từ Nuxt 4 gửi lên
}
