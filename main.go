package main

import (
	"fmt"
	"log"
	"QuickWork/handlers"
	"QuickWork/middlewares" // Import thư mục middleware vào
	"QuickWork/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Chuỗi kết nối MySQL chuẩn (Thay đổi user:password cho đúng cấu hình máy bạn)
	// Định dạng: user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:root@123@tcp(127.0.0.1:3306)/quickwork_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Không thể kết nối đến MySQL: ", err)
	}

	fmt.Println("🔌 Kết nối MySQL thành công!")

	// Tự động quét và sinh ra 3 bảng: users, students, businesses
	err = DB.AutoMigrate(&models.User{}, &models.Student{}, &models.Business{},&models.Job{},&models.Application{})
	if err != nil {
		log.Fatal("❌ Lỗi cấu trúc Migration: ", err)
	}

	fmt.Println("🎉 Đã tự động tạo các bảng dữ liệu thành công dưới MySQL!")
app := fiber.New()
	app.Use(logger.New())

	// Route Đăng ký tài khoản
	app.Post("/api/auth/register", handlers.HandleRegister(DB))

	app.Post("/api/auth/login", handlers.HandleLogin(DB))

	

app.Get("/api/users/me", middlewares.Protected(), func(c *fiber.Ctx) error {
		// Lấy thông tin mà chú bảo vệ đã bốc từ Token ra lúc nãy
		userID := c.Locals("user_id")
		role := c.Locals("role")

		return c.JSON(fiber.Map{
			"message": "🔓 Bạn đã vượt qua trạm kiểm soát bảo mật thành công!",
			"user_id": userID,
			"role":    role,
		})
	})
app.Put("/api/profile/student", 
	middlewares.Protected(), 
	middlewares.RequireRole("student"), 
	handlers.UpdateStudentProfile(DB),
)

// Chỉ DOANH NGHIỆP mới được vào cập nhật hồ sơ doanh nghiệp
app.Put("/api/profile/business", 
	middlewares.Protected(), 
	middlewares.RequireRole("business"), 
	handlers.UpdateBusinessProfile(DB),
)
app.Post("/api/jobs", 
	middlewares.Protected(), 
	middlewares.RequireRole("business"), 
	handlers.CreateJob(DB),
)
// 1. API xem Job công khai (Không cài Middleware -> Ai cũng vào được)
app.Get("/api/jobs", handlers.GetAvailableJobs(DB))

// 2. API Ứng tuyển (Bảo mật nghiêm ngặt, chỉ cho student vào)
app.Post("/api/jobs/apply", 
	middlewares.Protected(), 
	middlewares.RequireRole("student"), 
	handlers.ApplyJob(DB),
)
	log.Fatal(app.Listen(":3000"))
}
