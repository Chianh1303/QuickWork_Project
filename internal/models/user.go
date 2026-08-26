package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password  string    `gorm:"not null;type:varchar(255)" json:"-"`
	Role      string    `gorm:"type:varchar(20);not null" json:"role"` // admin, business, student
	Status         string     `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Balance        float64    `gorm:"type:decimal(15,2);default:0.00" json:"balance"` // Khớp với decimal của bạn
	FailedAttempts int        `gorm:"default:0" json:"failed_attempts"`
	LockedUntil    *time.Time `json:"locked_until,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`

	Student  *Student  `json:"student,omitempty"`
	Business *Business `json:"business,omitempty"`
}
