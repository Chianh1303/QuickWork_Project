package models

import "time"

// ⏰ Bảng Lịch sử Chấm công (Check-in / Check-out)
type Attendance struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	StudentID    uint    `gorm:"not null" json:"student_id"`
	StudentIDRef Student `gorm:"foreignKey:StudentID" json:"-"`
	JobID        uint    `gorm:"not null" json:"job_id"`
	JobRef       Job     `gorm:"foreignKey:JobID" json:"job,omitempty"`

	// Lưu thời gian thực tế bấm nút
	CheckInTime  *time.Time `json:"check_in_time"`
	CheckOutTime *time.Time `json:"check_out_time"`

	// Trạng thái ca làm: "working" (đang làm), "completed" (đã ra ca)
	Status string `gorm:"type:varchar(20);default:'working'" json:"status"`

	// Tự động lưu vết log hệ thống
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
