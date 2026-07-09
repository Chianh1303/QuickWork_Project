package models

import "time"

type Wallet struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;uniqueIndex" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Balance   float64   `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
