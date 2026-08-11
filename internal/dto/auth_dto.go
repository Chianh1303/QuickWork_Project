package dto

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	CompanyName string `json:"company_name"`
	TaxCode     string `json:"tax_code"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSummaryDTO struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type LoginResponse struct {
	Message string         `json:"message"`
	Token   string         `json:"token"`
	User    UserSummaryDTO `json:"user"`
}
