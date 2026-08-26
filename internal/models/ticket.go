package models

import "time"

type Ticket struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	ApplicationID uint       `gorm:"not null;uniqueIndex:idx_ticket_app_reporter" json:"application_id"`
	ReporterID    uint       `gorm:"not null;uniqueIndex:idx_ticket_app_reporter;index:idx_ticket_reporter" json:"reporter_id"` // User ID người khiếu nại
	TargetID      uint       `gorm:"not null;index:idx_ticket_target" json:"target_id"`                                           // User ID đối tượng bị khiếu nại
	Reason          string     `gorm:"type:varchar(255);not null" json:"reason"`
	Description     string     `gorm:"type:text;not null" json:"description"`
	RequestedAction string     `gorm:"type:varchar(255)" json:"requested_action,omitempty"`
	EvidenceURL     string     `gorm:"type:varchar(500)" json:"evidence_url,omitempty"`
	Status          string     `gorm:"type:varchar(20);default:'pending';index:idx_ticket_status" json:"status"` // pending, resolved, rejected
	Verdict         string     `gorm:"type:text" json:"verdict"`                                                 // Phán quyết của Admin
	ResolvedBy      *uint      `json:"resolved_by"`
	ResolvedAt      *time.Time `json:"resolved_at"`
	IsReappealed    bool       `gorm:"default:false" json:"is_reappealed"`
	ReappealReason  string     `gorm:"type:text" json:"reappeal_reason,omitempty"`
	ReappealEvidence string    `gorm:"type:varchar(500)" json:"reappeal_evidence,omitempty"`
	ReappealedAt    *time.Time `json:"reappealed_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`

	Application Application `gorm:"foreignKey:ApplicationID" json:"application,omitempty"`
	Reporter    User        `gorm:"foreignKey:ReporterID" json:"reporter,omitempty"`
	Target      User        `gorm:"foreignKey:TargetID" json:"target,omitempty"`
}
