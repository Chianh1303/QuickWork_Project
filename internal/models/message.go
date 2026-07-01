package models

import (
	"time"
)

type Message struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ApplicationID uint      `gorm:"index;not null" json:"application_id"`
	SenderID      uint      `gorm:"not null" json:"sender_id"` // UserID người gửi (Student hoặc Business)
	ReceiverID    uint      `gorm:"not null" json:"receiver_id"`
	MessageText   string    `gorm:"type:text;not null" json:"message_text"`
	IsRead        bool      `gorm:"default:false" json:"is_read"`
	CreatedAt     time.Time `json:"created_at"`
}
