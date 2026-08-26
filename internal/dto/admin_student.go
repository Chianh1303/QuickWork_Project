package dto

import "time"

type AdminStudentItem struct {
	StudentID uint      `json:"student_id"`
	UserID    uint      `json:"user_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	AvatarURL string    `json:"avatar_url"`
	Skills    string    `json:"skills"`
	CvURL     string    `json:"cv_url"`
	Status    string    `json:"status"` // approved, locked
	CreatedAt time.Time `json:"created_at"`
}

type AdminStudentListResponse struct {
	Items      []AdminStudentItem `json:"items"`
	Pagination PaginationMeta     `json:"pagination"`
}

type UpdateStudentStatusRequest struct {
	Status string `json:"status"` // approved, locked
}
