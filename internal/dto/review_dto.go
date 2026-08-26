package dto

type CreateReviewInput struct {
	ApplicationID uint   `json:"application_id"`
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
}
