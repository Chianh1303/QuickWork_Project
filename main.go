package main

import (
    "fmt"
    "log"
    "os"
    "QuickWork/handlers"
    "QuickWork/middlewares"
    "QuickWork/models"
    "QuickWork/database"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/cors" // Bổ sung CORS để gọi mượt từ Nuxt 4 (Port 3001)
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func main() {
    _ = os.MkdirAll("./uploads/avatars", os.ModePerm)
    _ = os.MkdirAll("./uploads/cvs", os.ModePerm)
    // 1. Khởi tạo chuỗi kết nối MySQL
    dsn := "root:root@123@tcp(127.0.0.1:3306)/quickwork_db?charset=utf8mb4&parseTime=True&loc=Local"
    
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Không thể kết nối đến MySQL: ", err)
    }
    fmt.Println("🔌 Kết nối MySQL thành công!")

    // 2. Chạy AutoMigrate để sinh cấu trúc bảng trước
    err = DB.AutoMigrate(&models.User{}, &models.Student{}, &models.Business{}, &models.Job{}, &models.Application{})
    if err != nil {
        log.Fatal("❌ Lỗi cấu trúc Migration: ", err)
    }
    fmt.Println("🎉 Đã tự động tạo các bảng dữ liệu thành công dưới MySQL!")

    // 3. Tiến hành Nạp Seed Data thực tế (Chỉ chạy sau khi bảng đã tồn tại an toàn)
    database.SeedDatabase(DB)

    // 4. Khởi tạo Fiber App
    app := fiber.New()
    app.Use(logger.New())

    // 5. Cấu hình CORS (BẮT BUỘC để tránh lỗi chặn liên cổng khi Nuxt 3001 gọi sang Go 3000)
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:3001",
        AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
        AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
        AllowCredentials: true,
    }))
    app.Static("/uploads", "./uploads")


    // ==========================================
    // CÁC ROUTE API HỆ THỐNG
    // ==========================================
    
    // Auth Routes
    app.Post("/api/auth/register", handlers.HandleRegister(DB))
    app.Post("/api/auth/login", handlers.HandleLogin(DB))

    // User Check Me
    app.Get("/api/users/me", middlewares.Protected(), func(c *fiber.Ctx) error {
        userID := c.Locals("user_id")
        role := c.Locals("role")

        return c.JSON(fiber.Map{
            "message": "🔓 Bạn đã vượt qua trạm kiểm soát bảo mật thành công!",
            "user_id": userID,
            "role":    role,
        })
    })

    
   // Profile Routes
    // --- Đối với Sinh viên ---
    app.Get("/api/profile/student", middlewares.Protected(), middlewares.RequireRole("student"), handlers.GetStudentProfile(DB))
    app.Put("/api/profile/student", middlewares.Protected(), middlewares.RequireRole("student"), handlers.UpdateStudentProfile(DB))

    // --- Đối với Doanh nghiệp ---
    app.Get("/api/profile/business", middlewares.Protected(), middlewares.RequireRole("business"), handlers.GetBusinessProfile(DB))
    app.Put("/api/profile/business", middlewares.Protected(), middlewares.RequireRole("business"), handlers.UpdateBusinessProfile(DB))
    
    // Jobs & Applications Routes
    app.Post("/api/jobs", middlewares.Protected(), middlewares.RequireRole("business"), handlers.CreateJob(DB))
    app.Get("/api/jobs", handlers.GetAvailableJobs(DB)) // Công khai
    
    app.Post("/api/jobs/apply", middlewares.Protected(), middlewares.RequireRole("student"), handlers.ApplyJob(DB))
    app.Put("/api/jobs/review-application", middlewares.Protected(), middlewares.RequireRole("business"), handlers.ReviewApplication(DB))


    // Sinh viên xem lịch sử ứng tuyển của mình (C6)
    app.Get("/api/applications/my-applications", 
        middlewares.Protected(), 
        middlewares.RequireRole("student"), 
        handlers.GetStudentApplications(DB),
    )

    // 🌟 THÊM ROUTE: Sinh viên thực hiện hủy ứng tuyển đơn còn chờ duyệt (C6)
    app.Post("/api/applications/:id/cancel",
        middlewares.Protected(),
        middlewares.RequireRole("student"),
        handlers.CancelApplication(DB),
    )

    // Doanh nghiệp xem danh sách ứng viên nộp đơn (B6)
    app.Get("/api/applications/employer", 
        middlewares.Protected(), 
        middlewares.RequireRole("business"), 
        handlers.GetEmployerApplications(DB),
    )
    
	// Đặt cùng nhóm với các route của Student nhé Chanh
app.Post("/api/applications/respond-offer", 
    middlewares.Protected(), 
    middlewares.RequireRole("student"), 
    handlers.RespondToOffer(DB),
)
    // Khởi động Server tại cổng 3000
    log.Fatal(app.Listen(":3000"))
}