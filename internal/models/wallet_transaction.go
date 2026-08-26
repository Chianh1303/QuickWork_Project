package models

import "time"

type WalletTransaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	WalletID      uint      `gorm:"not null;index" json:"wallet_id"`
	Wallet        Wallet    `gorm:"foreignKey:WalletID" json:"wallet,omitempty"`
	Type          string    `gorm:"type:varchar(30);not null" json:"type"`
	Amount        float64   `gorm:"type:decimal(15,2);not null" json:"amount"`
	Description   string    `gorm:"type:varchar(255)" json:"description"`
	ReferenceID   uint      `json:"reference_id"`
	ReferenceType string    `gorm:"type:varchar(50)" json:"reference_type"`
	CreatedAt     time.Time `json:"created_at"`
}
