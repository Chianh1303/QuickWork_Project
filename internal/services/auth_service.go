package services

import (
	"errors"
	"fmt"
	"time"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	jwtSecret = []byte("quickwork_secret_key_2026")

	ErrEmptyFields           = errors.New("Email, password và role không được để trống")
	ErrInvalidRole           = errors.New("Role phải là 'student' hoặc 'business'")
	ErrEmailExists           = errors.New("Email này đã được đăng ký sử dụng")
	ErrTaxCodeExists         = errors.New("Mã số thuế này đã được đăng ký sử dụng")
	ErrStudentNameRequired   = errors.New("Họ và tên sinh viên không được để trống")
	ErrBusinessFieldsRequired = errors.New("Tên công ty và mã số thuế không được để trống")
	ErrInvalidCredentials    = errors.New("Email hoặc mật khẩu không chính xác")
)

type AccountForbiddenError struct {
	Message string
	Status  string
}

func (e *AccountForbiddenError) Error() string {
	return e.Message
}

type AccountLockedError struct {
	Message          string
	RemainingSeconds int
}

func (e *AccountLockedError) Error() string {
	return e.Message
}

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	if req.Email == "" || req.Password == "" || req.Role == "" {
		return ErrEmptyFields
	}
	if req.Role != "student" && req.Role != "business" {
		return ErrInvalidRole
	}

	existingUser, err := s.authRepo.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Lỗi mã hóa hệ thống")
	}

	status := "approved"
	if req.Role == "business" {
		status = "pending"
	}

	newUser := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
		Status:   status,
		Balance:  0.0,
	}

	var studentProfile *models.Student
	var businessProfile *models.Business

	if req.Role == "student" {
		if req.FullName == "" {
			return ErrStudentNameRequired
		}
		studentProfile = &models.Student{
			FullName: req.FullName,
			Phone:    req.Phone,
		}
	} else if req.Role == "business" {
		if req.CompanyName == "" || req.TaxCode == "" {
			return ErrBusinessFieldsRequired
		}
		existingBiz, err := s.authRepo.GetBusinessByTaxCode(req.TaxCode)
		if err == nil && existingBiz != nil {
			return ErrTaxCodeExists
		}
		businessProfile = &models.Business{
			CompanyName: req.CompanyName,
			TaxCode:     req.TaxCode,
			Phone:       req.Phone,
		}
	}

	return s.authRepo.CreateUserWithProfile(newUser, studentProfile, businessProfile)
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, ErrInvalidCredentials
	}

	user, err := s.authRepo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// 1. Kiểm tra khóa theo thời gian (Progressive Lockout Check)
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		remainingSec := int(time.Until(*user.LockedUntil).Seconds()) + 1
		return nil, &AccountLockedError{
			Message:          fmt.Sprintf("⚠️ Tài khoản tạm bị khóa do nhập sai mật khẩu nhiều lần. Vui lòng thử lại sau %d giây!", remainingSec),
			RemainingSeconds: remainingSec,
		}
	}

	// 2. Kiểm tra mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		user.FailedAttempts++
		now := time.Now()

		if user.FailedAttempts >= 10 {
			user.Status = "suspended"
			user.LockedUntil = nil
			_ = s.authRepo.UpdateUser(user)
			return nil, &AccountForbiddenError{
				Message: "🛑 Tài khoản của bạn đã bị tạm ngưng hoạt động do nhập sai mật khẩu quá 10 lần.",
				Status:  "suspended",
			}
		} else if user.FailedAttempts >= 8 {
			lockedTime := now.Add(15 * time.Minute)
			user.LockedUntil = &lockedTime
			_ = s.authRepo.UpdateUser(user)
			return nil, &AccountLockedError{
				Message:          "⚠️ Nhập sai 8 lần liên tiếp. Tài khoản bị tạm khóa 15 phút!",
				RemainingSeconds: 900,
			}
		} else if user.FailedAttempts >= 5 {
			lockedTime := now.Add(1 * time.Minute)
			user.LockedUntil = &lockedTime
			_ = s.authRepo.UpdateUser(user)
			return nil, &AccountLockedError{
				Message:          "⚠️ Nhập sai 5 lần liên tiếp. Tài khoản bị tạm khóa 1 phút!",
				RemainingSeconds: 60,
			}
		}

		_ = s.authRepo.UpdateUser(user)
		return nil, ErrInvalidCredentials
	}

	// 3. Mật khẩu chính xác -> Reset số lần sai về 0
	if user.FailedAttempts > 0 || user.LockedUntil != nil {
		user.FailedAttempts = 0
		user.LockedUntil = nil
		_ = s.authRepo.UpdateUser(user)
	}

	if user.Status != "approved" && user.Status != "active" {
		return nil, &AccountForbiddenError{
			Message: "Tài khoản của bạn chưa được duyệt hoặc đang bị khóa",
			Status:  user.Status,
		}
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("Không thể tạo token bảo mật")
	}

	return &dto.LoginResponse{
		Message: "🔒 Đăng nhập thành công!",
		Token:   tokenString,
		User: dto.UserSummaryDTO{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
