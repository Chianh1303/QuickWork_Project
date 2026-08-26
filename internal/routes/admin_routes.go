package routes

import (
	"QuickWork/internal/controllers"
	"QuickWork/internal/middleware"
	"QuickWork/internal/repositories"
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(app *fiber.App, db *gorm.DB) {
	adminRepo := repositories.NewAdminRepository(db)
	notifRepo := repositories.NewNotificationRepository(db)
	notifService := services.NewNotificationService(notifRepo)
	adminService := services.NewAdminService(adminRepo, notifService)
	adminController := controllers.NewAdminController(adminService)

	admin := app.Group("/api/admin", middleware.Protected(), middleware.RequireRole("admin"))

	admin.Get("/dashboard/stats", adminController.GetAdminDashboardStats)
	admin.Get("/businesses/pending", adminController.GetPendingBusinesses)
	admin.Get("/businesses/:id", adminController.GetBusinessKYBDetail)
	admin.Put("/businesses/:id/review", adminController.ReviewBusinessKYB)

	// Quản lý Sinh viên (Feature A6)
	admin.Get("/students", adminController.GetStudents)
	admin.Get("/students/:id", adminController.GetStudentDetail)
	admin.Put("/students/:id/status", adminController.UpdateStudentStatus)

	// Quản lý Doanh nghiệp (Feature A7)
	admin.Get("/businesses", adminController.GetBusinesses)
	admin.Put("/businesses/:id/status", adminController.UpdateBusinessStatus)

	// Xử lý Khiếu nại (Feature A8)
	admin.Get("/tickets", adminController.GetTickets)
	admin.Get("/tickets/:id", adminController.GetTicketDetail)
	admin.Put("/tickets/:id/resolve", adminController.ResolveTicket)

	// Quản lý Danh mục & Kỹ năng (Feature A9)
	admin.Get("/categories", adminController.GetCategories)
	admin.Post("/categories", adminController.CreateCategory)
	admin.Delete("/categories/:id", adminController.DeleteCategory)
	admin.Get("/skills", adminController.GetSkills)
	admin.Post("/skills", adminController.CreateSkill)
	admin.Delete("/skills/:id", adminController.DeleteSkill)

	// Duyệt Bài Tuyển Dụng (Admin Job Approval)
	admin.Get("/jobs/pending", adminController.GetPendingJobs)
	admin.Put("/jobs/:id/status", adminController.UpdateJobStatus)
}
