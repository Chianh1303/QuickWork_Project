package models

import "time"

type Notification struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"index;not null" json:"user_id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Message     string    `gorm:"type:text;not null" json:"message"`
	Type        string    `gorm:"type:varchar(50);default:'system'" json:"type"` // chat, application, offer, escrow, system
	ReferenceID uint      `gorm:"default:0" json:"reference_id"`
	IsRead      bool      `gorm:"default:false" json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}
