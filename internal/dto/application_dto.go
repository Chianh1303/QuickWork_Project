package dto

type ApplyJobInput struct {
	JobID     uint   `json:"job_id"`
	CoverNote string `json:"cover_note"`
}

type ReviewApplicationInput struct {
	ApplicationID  uint   `json:"application_id"`
	Status         string `json:"status"`
	OfferSalary    string `json:"offer_salary"`
	OfferStartDate string `json:"offer_start_date"`
	OfferMessage   string `json:"offer_message"`
}

type RespondOfferInput struct {
	ApplicationID uint   `json:"application_id"`
	Response      string `json:"response"`
}

type CompleteJobInput struct {
	ApplicationID      uint   `json:"application_id"`
	CompletionNote     string `json:"completion_note"`
	CompletionProofUrl string `json:"completion_proof_url"`
}
