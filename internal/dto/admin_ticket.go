package dto

import "time"

type AdminTicketItem struct {
	TicketID      uint       `json:"ticket_id"`
	ApplicationID uint       `json:"application_id"`
	ReporterID    uint       `json:"reporter_id"`
	ReporterEmail string     `json:"reporter_email"`
	ReporterRole  string     `json:"reporter_role"`
	TargetID      uint       `json:"target_id"`
	TargetEmail   string     `json:"target_email"`
	Reason        string     `json:"reason"`
	Description   string     `json:"description"`
	Status        string     `json:"status"` // pending, resolved, rejected
	Verdict       string     `json:"verdict"`
	ResolvedAt    *time.Time `json:"resolved_at"`
	CreatedAt     time.Time  `json:"created_at"`
}

type AdminTicketListResponse struct {
	Items      []AdminTicketItem `json:"items"`
	Pagination PaginationMeta    `json:"pagination"`
}

type ResolveTicketRequest struct {
	Verdict string `json:"verdict"`
	Status  string `json:"status"` // resolved, rejected
}
