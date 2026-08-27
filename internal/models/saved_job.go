package models

import "time"

type SavedJob struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StudentID uint      `gorm:"not null;index:idx_student_job,unique" json:"student_id"`
	JobID     uint      `gorm:"not null;index:idx_student_job,unique" json:"job_id"`
	Job       Job       `gorm:"foreignKey:JobID" json:"job,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
