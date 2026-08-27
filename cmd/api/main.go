package main

import (
	_ "QuickWork/docs"
	"QuickWork/internal/config"
	"QuickWork/internal/database"
	"QuickWork/internal/routes"
	"QuickWork/internal/storage"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Bổ sung CORS để gọi mượt từ Nuxt (Port 3001)
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			if os.Getenv(key) == "" {
				_ = os.Setenv(key, val)
			}
		}
	}
}

func main() {
	loadEnv()
	_ = os.MkdirAll(config.AvatarDir, os.ModePerm)
	_ = os.MkdirAll(config.CVDir, os.ModePerm)

	var err error
	DB, err = database.Connect(config.MySQLDSN)
	if err != nil {
		log.Fatal("❌ Không thể kết nối đến MySQL: ", err)
	}
	fmt.Println("🔌 Kết nối MySQL thành công!")

	// 2. Chạy AutoMigrate và Seed Data không làm nản luồng lắng nghe Cổng Server
	go func() {
		if err := database.AutoMigrate(DB); err != nil {
			log.Println("⚠️ Lỗi cấu trúc Migration: ", err)
			return
		}
		fmt.Println("🎉 Đã tự động tạo các bảng dữ liệu thành công dưới MySQL!")
		database.SeedDatabase(DB)
	}()

	// 4. Khởi tạo AWS S3 / Local Storage Engine
	_ = storage.NewStorageProvider()

	// 4. Khởi tạo Fiber App
	app := fiber.New()
	app.Use(logger.New())

	// 5. Cấu hình CORS (BẮT BUỘC để cho phép Frontend Nuxt gọi API từ bất kỳ tên miền nào)
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))
	app.Static("/uploads", config.UploadDir)

	// 6. Cấu hình Rate Limiting bảo vệ các Route nhạy cảm (Chống Brute Force & AI Spam)
	app.Use("/api/auth/login", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "⚠️ Bạn đã thử đăng nhập quá 10 lần trong 1 phút. Vui lòng đợi 1 phút!",
			})
		},
	}))

	app.Use("/api/ai/*", limiter.New(limiter.Config{
		Max:        15,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "⚠️ Bạn đã gọi dịch vụ AI quá 15 lần trong 1 phút. Vui lòng đợi 1 phút!",
			})
		},
	}))

	// Swagger OpenAPI UI Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health Check & Root Welcome Endpoints (Render Health Checker & Uptime Monitor)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "⚡ QuickWork Backend API Server is running smoothly!",
			"service": "QuickWork API Core Engine",
			"time":    time.Now().Format(time.RFC3339),
		})
	})
	app.Head("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	// ==========================================
	// CÁC ROUTE API HỆ THỐNG
	// ==========================================
	routes.Register(app, DB)

	// Khởi động Server tại cổng 3000
	log.Fatal(app.Listen(config.ServerAddr))

}
