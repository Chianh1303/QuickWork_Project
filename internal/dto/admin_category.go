package dto

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSkillRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
