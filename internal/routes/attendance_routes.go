package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAttendanceRoutes(app *fiber.App, db *gorm.DB) {
	// ⏰ 2. Tuyển tập Route Chấm công (Check-in / Check-out) viết theo chuẩn Fiber của Chanh
	app.Post("/api/attendance/check-in",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.CheckIn(db),
	)

	app.Post("/api/attendance/check-out",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.CheckOut(db),
	)
}
