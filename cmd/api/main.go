package main

import (
	_ "QuickWork/docs"
	"QuickWork/internal/config"
	"QuickWork/internal/database"
	"QuickWork/internal/queue"
	"QuickWork/internal/routes"
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

	// 2. Chạy AutoMigrate để sinh cấu trúc bảng trước
	err = database.AutoMigrate(DB)
	if err != nil {
		log.Fatal("❌ Lỗi cấu trúc Migration: ", err)
	}
	fmt.Println("🎉 Đã tự động tạo các bảng dữ liệu thành công dưới MySQL!")

	// 3. Tiến hành Nạp Seed Data thực tế (Chỉ chạy sau khi bảng đã tồn tại an toàn)
	database.SeedDatabase(DB)

	// 4. Khởi tạo RabbitMQ Client & Đăng ký Background Workers
	rmqClient := queue.NewRabbitMQClient(config.RabbitMQURL)
	defer rmqClient.Close()
	queue.RegisterWorkers(rmqClient)

	// 5. Khởi tạo Fiber App
	app := fiber.New()
	app.Use(logger.New())

	// 5. Cấu hình CORS (BẮT BUỘC để tránh lỗi chặn liên cổng khi Nuxt 3001 gọi sang Go 3000)
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.CORSOrigins,

		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",

		AllowHeaders: "Origin,Content-Type,Accept,Authorization",

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

	// ==========================================
	// CÁC ROUTE API HỆ THỐNG
	// ==========================================
	routes.Register(app, DB)

	// Khởi động Server tại cổng 3000
	log.Fatal(app.Listen(config.ServerAddr))

}
