package dto

type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	CompanyName string `json:"company_name"`
	TaxCode     string `json:"tax_code"`
	Address     string `json:"address"`
	Website     string `json:"website"`
	CompanySize string `json:"company_size"`
	Description string `json:"description"`
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

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OTPCode     string `json:"otp_code"`
	NewPassword string `json:"new_password"`
}

type GoogleLoginRequest struct {
	IDToken string `json:"id_token"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Role    string `json:"role"`
}
