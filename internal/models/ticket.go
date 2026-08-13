package models

import "time"

type Ticket struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	ApplicationID uint       `gorm:"not null" json:"application_id"`
	ReporterID    uint       `gorm:"not null" json:"reporter_id"` // User ID người khiếu nại
	TargetID      uint       `gorm:"not null" json:"target_id"`   // User ID đối tượng bị khiếu nại
	Reason        string     `gorm:"type:varchar(255);not null" json:"reason"`
	Description   string     `gorm:"type:text;not null" json:"description"`
	Status        string     `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, resolved, rejected
	Verdict       string     `gorm:"type:text" json:"verdict"`                         // Phán quyết của Admin
	ResolvedBy    *uint      `json:"resolved_by"`
	ResolvedAt    *time.Time `json:"resolved_at"`
	CreatedAt     time.Time  `json:"created_at"`

	Application Application `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	Reporter    User        `gorm:"foreignKey:ReporterID" json:"reporter,omitempty"`
	Target      User        `gorm:"foreignKey:TargetID" json:"target,omitempty"`
}
