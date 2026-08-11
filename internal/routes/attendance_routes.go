package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAttendanceRoutes(app *fiber.App, db *gorm.DB) {
	attendanceRepo := repositories.NewAttendanceRepository(db)
	attendanceService := services.NewAttendanceService(attendanceRepo)
	attendanceController := controllers.NewAttendanceController(attendanceService)

	app.Post("/api/attendance/check-in",
		middleware.Protected(),
		middleware.RequireRole("student"),
		attendanceController.CheckIn,
	)

	app.Post("/api/attendance/check-out",
		middleware.Protected(),
		middleware.RequireRole("student"),
		attendanceController.CheckOut,
	)
}
