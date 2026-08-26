package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterApplicationRoutes(app *fiber.App, db *gorm.DB) {
	notifRepo := repositories.NewNotificationRepository(db)
	notifService := services.NewNotificationService(notifRepo)
	appRepo := repositories.NewApplicationRepository(db)
	appService := services.NewApplicationService(appRepo, notifService)
	appController := controllers.NewApplicationController(appService)

	app.Post("/api/jobs/apply", middleware.Protected(), middleware.RequireRole("student"), appController.ApplyJob)

	// Sinh viên xem lịch sử ứng tuyển của mình (C6)
	app.Get("/api/applications/my-applications",
		middleware.Protected(),
		middleware.RequireRole("student"),
		appController.GetStudentApplications,
	)

	// 🌟 THÊM ROUTE: Sinh viên thực hiện hủy ứng tuyển đơn còn chờ duyệt (C6)
	app.Post("/api/applications/:id/cancel",
		middleware.Protected(),
		middleware.RequireRole("student"),
		appController.CancelApplication,
	)

	// Doanh nghiệp xem danh sách ứng viên nộp đơn (B6)
	app.Get("/api/applications/employer",
		middleware.Protected(),
		middleware.RequireRole("business"),
		appController.GetEmployerApplications,
	)

	app.Post("/api/applications/student-complete",
		middleware.Protected(),
		middleware.RequireRole("student"),
		appController.StudentCompleteJob,
	)

	app.Post("/api/applications/business-complete",
		middleware.Protected(),
		middleware.RequireRole("business"),
		appController.BusinessCompleteJob,
	)
}
