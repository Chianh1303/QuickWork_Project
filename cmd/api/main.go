package main

import (
	"QuickWork/internal/config"
	"QuickWork/internal/database"
	"QuickWork/internal/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Bổ sung CORS để gọi mượt từ Nuxt 4 (Port 3001)
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
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

	// 4. Khởi tạo Fiber App
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

	// ==========================================
	// CÁC ROUTE API HỆ THỐNG
	// ==========================================
	routes.Register(app, DB)

	// Khởi động Server tại cổng 3000
	log.Fatal(app.Listen(config.ServerAddr))

}
