package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"QuickWork/internal/cache"
	"QuickWork/internal/dto"
	"QuickWork/internal/models"
	"QuickWork/internal/queue"
	"QuickWork/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	jwtSecret = []byte("quickwork_secret_key_2026")

	ErrEmptyFields            = errors.New("Email, password và role không được để trống")
	ErrInvalidRole            = errors.New("Role phải là 'student' hoặc 'business'")
	ErrEmailExists            = errors.New("Email này đã được đăng ký sử dụng")
	ErrTaxCodeExists          = errors.New("Mã số thuế này đã được đăng ký sử dụng")
	ErrStudentNameRequired    = errors.New("Họ và tên sinh viên không được để trống")
	ErrBusinessFieldsRequired = errors.New("Tên công ty và mã số thuế không được để trống")
	ErrInvalidCredentials     = errors.New("Email hoặc mật khẩu không chính xác")
	ErrUserNotFound           = errors.New("Tài khoản với Email này không tồn tại trên hệ thống")
	ErrInvalidOTP             = errors.New("Mã OTP không chính xác hoặc đã hết hạn 5 phút")
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

// In-memory OTP storage fallback when Redis is not available
var (
	otpMemoryStore = make(map[string]otpItem)
	otpStoreMutex  sync.RWMutex
)

type otpItem struct {
	Code      string
	ExpiresAt time.Time
}

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	SendPasswordResetOTP(email string) error
	VerifyPasswordResetOTP(email string, otpCode string, newPassword string) error
	GoogleLogin(req dto.GoogleLoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	authRepo     repositories.AuthRepository
	emailService EmailService
	cacheClient  cache.CacheClient
	rmqClient    queue.RabbitMQClient
}

func NewAuthService(authRepo repositories.AuthRepository, emailSvc EmailService, cacheClient cache.CacheClient, rmqClient queue.RabbitMQClient) AuthService {
	return &authService{
		authRepo:     authRepo,
		emailService: emailSvc,
		cacheClient:  cacheClient,
		rmqClient:    rmqClient,
	}
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

	return s.generateAuthResponse(user, "🔒 Đăng nhập thành công!")
}

func (s *authService) SendPasswordResetOTP(email string) error {
	if email == "" {
		return errors.New("Email không được để trống")
	}

	user, err := s.authRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return ErrUserNotFound
	}

	// Generate 6-digit random numeric OTP code
	nBig, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return errors.New("Lỗi tạo mã OTP hệ thống")
	}
	otpCode := fmt.Sprintf("%06d", nBig.Int64()+100000)

	// Save OTP to Redis (5 min TTL) or In-Memory fallback
	otpKey := fmt.Sprintf("otp_reset:%s", email)
	if s.cacheClient != nil && s.cacheClient.IsAvailable() {
		_ = s.cacheClient.Set(context.Background(), otpKey, otpCode, 5*time.Minute)
	} else {
		otpStoreMutex.Lock()
		otpMemoryStore[email] = otpItem{
			Code:      otpCode,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		}
		otpStoreMutex.Unlock()
	}

	// Off-thread Async Queueing via RabbitMQ or Direct Service Call
	if s.rmqClient != nil && s.rmqClient.IsAvailable() {
		_ = s.rmqClient.Publish(queue.QueueEmailOTP, queue.EmailOTPPayload{
			Email:   email,
			OTPCode: otpCode,
		})
	} else if s.emailService != nil {
		go func() {
			_ = s.emailService.SendOTPEmail(email, otpCode)
		}()
	}

	return nil
}

func (s *authService) VerifyPasswordResetOTP(email string, otpCode string, newPassword string) error {
	if email == "" || otpCode == "" || newPassword == "" {
		return errors.New("Email, mã OTP và mật khẩu mới không được để trống")
	}

	user, err := s.authRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return ErrUserNotFound
	}

	// Verify OTP Code
	otpKey := fmt.Sprintf("otp_reset:%s", email)
	var isValid bool

	if s.cacheClient != nil && s.cacheClient.IsAvailable() {
		storedVal, err := s.cacheClient.Get(context.Background(), otpKey)
		if err == nil && storedVal == otpCode {
			isValid = true
			_ = s.cacheClient.Delete(context.Background(), otpKey)
		}
	} else {
		otpStoreMutex.Lock()
		item, exists := otpMemoryStore[email]
		if exists && item.Code == otpCode && time.Now().Before(item.ExpiresAt) {
			isValid = true
			delete(otpMemoryStore, email)
		}
		otpStoreMutex.Unlock()
	}

	if !isValid {
		return ErrInvalidOTP
	}

	// Hash new password and unlock account
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Lỗi mã hóa mật khẩu mới")
	}

	user.Password = string(hashedPassword)
	user.FailedAttempts = 0
	user.LockedUntil = nil
	if user.Status == "suspended" || user.Status == "locked" {
		user.Status = "approved"
	}

	return s.authRepo.UpdateUser(user)
}

func (s *authService) GoogleLogin(req dto.GoogleLoginRequest) (*dto.LoginResponse, error) {
	if req.Email == "" {
		return nil, errors.New("Email Google không hợp lệ")
	}

	user, err := s.authRepo.GetUserByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
		// Auto-register new Google user
		role := req.Role
		if role != "business" {
			role = "student"
		}

		dummyPassword, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("google_%d", time.Now().UnixNano())), bcrypt.DefaultCost)

		newUser := &models.User{
			Email:    req.Email,
			Password: string(dummyPassword),
			Role:     role,
			Status:   "approved",
			Balance:  0.0,
		}

		var studentProfile *models.Student
		var businessProfile *models.Business

		fullName := req.Name
		if fullName == "" {
			fullName = strings.Split(req.Email, "@")[0]
		}

		if role == "student" {
			studentProfile = &models.Student{
				FullName:  fullName,
				AvatarUrl: req.Picture,
			}
		} else {
			businessProfile = &models.Business{
				CompanyName: fullName,
				TaxCode:     fmt.Sprintf("GOOG_%d", time.Now().UnixNano()%10000000),
				LogoUrl:     req.Picture,
			}
		}

		if err := s.authRepo.CreateUserWithProfile(newUser, studentProfile, businessProfile); err != nil {
			return nil, fmt.Errorf("Không thể tạo tài khoản Google: %w", err)
		}
		user = newUser
	}

	if user.Status != "approved" && user.Status != "active" {
		return nil, &AccountForbiddenError{
			Message: "Tài khoản của bạn đang bị khóa hoặc ngưng hoạt động",
			Status:  user.Status,
		}
	}

	return s.generateAuthResponse(user, "🚀 Đăng nhập bằng Google thành công!")
}

func (s *authService) generateAuthResponse(user *models.User, message string) (*dto.LoginResponse, error) {
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
		Message: message,
		Token:   tokenString,
		User: dto.UserSummaryDTO{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
