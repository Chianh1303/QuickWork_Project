package handlers

import (
	"QuickWork/internal/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Cấu trúc nhận dữ liệu Đăng ký từ Client gửi lên
type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Role        string `json:"role"`         // student hoặc business
	FullName    string `json:"full_name"`    // Cho Student
	Phone       string `json:"phone"`        // Cho cả 2
	CompanyName string `json:"company_name"` // Cho Business
	TaxCode     string `json:"tax_code"`     // Cho Business
}

func HandleRegister(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		// 1. Kiểm tra dữ liệu bắt buộc cơ bản
		if req.Email == "" || req.Password == "" || req.Role == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email, password và role không được để trống"})
		}
		if req.Role != "student" && req.Role != "business" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Role phải là 'student' hoặc 'business'"})
		}

		// 2. Kiểm tra xem Email đã tồn tại trong DB chưa
		var existingUser models.User
		err := db.Where("email = ?", req.Email).First(&existingUser).Error
		if err == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Email này đã được đăng ký sử dụng"})
		}

		// 3. Mã hóa mật khẩu (Hash Password) bằng bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Lỗi mã hóa hệ thống"})
		}

		// 4. Dùng Transaction để lưu đồng thời vào các bảng liên quan
		err = db.Transaction(func(tx *gorm.DB) error {
			status := "approved"
			if req.Role == "business" {
				// Doanh nghiệp đăng ký phải chờ Admin duyệt KYB, nên chưa được cấp quyền đăng nhập dashboard ngay.
				status = "pending"
			}

			// Tạo User gốc trước
			newUser := models.User{
				Email:    req.Email,
				Password: string(hashedPassword),
				Role:     req.Role,
				Status:   status,
				Balance:  0.0,
			}
			if err := tx.Create(&newUser).Error; err != nil {
				return err // Lỗi tự động Rollback
			}

			// Tùy theo Role mà tạo tiếp bảng hồ sơ tương ứng
			if req.Role == "student" {
				if req.FullName == "" {
					return fmt.Errorf("họ và tên sinh viên không được để trống")
				}
				studentProfile := models.Student{
					UserID:   newUser.ID, // Lấy ID vừa sinh ra ở trên gán vào khóa ngoại
					FullName: req.FullName,
					Phone:    req.Phone,
				}
				if err := tx.Create(&studentProfile).Error; err != nil {
					return err
				}
			} else if req.Role == "business" {
				if req.CompanyName == "" || req.TaxCode == "" {
					return fmt.Errorf("tên công ty và mã số thuế không được để trống")
				}
				businessProfile := models.Business{
					UserID:      newUser.ID,
					CompanyName: req.CompanyName,
					TaxCode:     req.TaxCode,
					Phone:       req.Phone,
				}
				if err := tx.Create(&businessProfile).Error; err != nil {
					return err
				}
			}

			return nil // Hoàn thành trơn tru -> Tự động Commit xuống DB
		})

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "🎉 Đăng ký tài khoản thành công!",
		})
	}
}

// Định nghĩa một chuỗi Secret Key bí mật để ký token (Trong thực tế sẽ giấu vào file .env)
var jwtSecret = []byte("quickwork_secret_key_2026")

// Cấu trúc nhận dữ liệu Đăng nhập từ Client
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		// 1. Tìm User theo Email dưới DB
		var user models.User
		err := db.Where("email = ?", req.Email).First(&user).Error
		if err != nil {
			// Để bảo mật, không báo rõ là "Sai email" mà báo chung chung là sai tài khoản/mật khẩu
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Email hoặc mật khẩu không chính xác"})
		}

		// 2. Kiểm tra mật khẩu (So sánh mật khẩu thuần với chuỗi Hash dưới DB)
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Email hoặc mật khẩu không chính xác"})
		}

		// Luôn kiểm tra password trước status để không làm lộ email nào đang tồn tại trong hệ thống.
		if user.Status != "approved" && user.Status != "active" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Tài khoản của bạn chưa được duyệt hoặc đang bị khóa",
				"status":  user.Status,
			})
		}

		// 3. Nếu mọi thứ đúng -> Tiến hành tạo Token JWT
		// Khai báo các thông tin chứa trong Token (Claims)
		claims := jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token có hạn trong 3 ngày
		}

		// Ký mã hóa token bằng thuật toán HS256 và Secret Key của mình
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Không thể tạo token bảo mật"})
		}

		// 4. Trả về Token kèm thông tin cơ bản cho Client
		return c.JSON(fiber.Map{
			"message": "🔒 Đăng nhập thành công!",
			"token":   tokenString,
			"user": fiber.Map{
				"id":    user.ID,
				"email": user.Email,
				"role":  user.Role,
			},
		})
	}
}
