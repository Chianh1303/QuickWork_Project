package dto

import "time"

type AdminTicketItem struct {
	TicketID         uint       `json:"ticket_id"`
	ApplicationID    uint       `json:"application_id"`
	ReporterID       uint       `json:"reporter_id"`
	ReporterEmail    string     `json:"reporter_email"`
	ReporterRole     string     `json:"reporter_role"`
	TargetID         uint       `json:"target_id"`
	TargetEmail      string     `json:"target_email"`
	Reason           string     `json:"reason"`
	Description      string     `json:"description"`
	RequestedAction  string     `json:"requested_action"`
	EvidenceURL      string     `json:"evidence_url"`
	Status           string     `json:"status"` // pending, resolved, rejected
	Verdict          string     `json:"verdict"`
	ResolvedAt       *time.Time `json:"resolved_at"`
	IsReappealed     bool       `json:"is_reappealed"`
	ReappealReason   string     `json:"reappeal_reason"`
	ReappealEvidence string     `json:"reappeal_evidence"`
	ReappealedAt     *time.Time `json:"reappealed_at"`
	CreatedAt        time.Time  `json:"created_at"`
}

type AdminTicketListResponse struct {
	Items      []AdminTicketItem `json:"items"`
	Pagination PaginationMeta    `json:"pagination"`
}

type ResolveTicketRequest struct {
	Verdict string `json:"verdict"`
	Status  string `json:"status"` // resolved, rejected
}

type CreateTicketRequest struct {
	ApplicationID   uint   `json:"application_id"`
	Reason          string `json:"reason"`
	Description     string `json:"description"`
	RequestedAction string `json:"requested_action"`
	EvidenceURL     string `json:"evidence_url"`
}

type ReappealTicketRequest struct {
	Reason   string `json:"reason"`
	Evidence string `json:"evidence"`
}
