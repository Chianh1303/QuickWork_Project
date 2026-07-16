package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApplicationRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/api/jobs/apply", middleware.Protected(), middleware.RequireRole("student"), handlers.ApplyJob(db))

	// Sinh viên xem lịch sử ứng tuyển của mình (C6)
	app.Get("/api/applications/my-applications",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.GetStudentApplications(db),
	)

	// 🌟 THÊM ROUTE: Sinh viên thực hiện hủy ứng tuyển đơn còn chờ duyệt (C6)
	app.Post("/api/applications/:id/cancel",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.CancelApplication(db),
	)

	// Doanh nghiệp xem danh sách ứng viên nộp đơn (B6)
	app.Get("/api/applications/employer",
		middleware.Protected(),
		middleware.RequireRole("business"),
		handlers.GetEmployerApplications(db),
	)

	app.Post("/api/applications/student-complete",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.StudentCompleteJob(db),
	)

	app.Post("/api/applications/business-complete",
		middleware.Protected(),
		middleware.RequireRole("business"),
		handlers.BusinessCompleteJob(db),
	)
}