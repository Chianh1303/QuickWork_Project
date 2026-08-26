package models

import "time"

type Review struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ApplicationID uint `json:"application_id"`

	ReviewerID uint `json:"reviewer_id"`

	RevieweeID uint `json:"reviewee_id"`

	Rating int `gorm:"not null" json:"rating"`

	Comment string `gorm:"type:text" json:"comment"`

	ReviewType string `gorm:"type:varchar(20)" json:"review_type"`

	CreatedAt time.Time `json:"created_at"`
}
