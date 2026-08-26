package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("quickwork_secret_key_2026")

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := strings.TrimSpace(c.Get("Authorization"))
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Header không hợp lệ"})
		}
		tokenString := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") || strings.HasPrefix(authHeader, "bearer ") {
			tokenString = strings.TrimSpace(authHeader[7:])
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "❌ Token không hợp lệ hoặc đã hết hạn"})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token không hợp lệ"})
		}
		c.Locals("user_id", claims["user_id"])
		c.Locals("role", claims["role"])
		return c.Next()
	}
}
func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Lấy role đã được hàm Protected() bốc ra trước đó
		userRole := c.Locals("role")
		if userRole == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "❌ Không tìm thấy thông tin phân quyền",
			})
		}

		// 2. Ép kiểu về string để so sánh
		roleStr, ok := userRole.(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "❌ Lỗi định dạng phân quyền",
			})
		}

		// 3. Kiểm tra xem role của user có nằm trong danh sách được phép không
		for _, role := range allowedRoles {
			if roleStr == role {
				return c.Next() // Đúng role -> Cho đi tiếp vào API nghiệp vụ
			}
		}

		// 4. Nếu không trùng khớp -> Chặn lại báo lỗi 403 Forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "🚫 Quyền truy cập bị từ chối. Bạn không có thẩm quyền thực hiện hành động này!",
		})
	}
}
