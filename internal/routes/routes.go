package routes

import (
	"QuickWork/internal/handlers"
	"QuickWork/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, db *gorm.DB) {
	// Auth Routes
	app.Post("/api/auth/register", handlers.HandleRegister(db))
	app.Post("/api/auth/login", handlers.HandleLogin(db))

	// User Check Me
	app.Get("/api/users/me", middleware.Protected(), func(c *fiber.Ctx) error {
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
	app.Get("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), handlers.GetStudentProfile(db))
	app.Put("/api/profile/student", middleware.Protected(), middleware.RequireRole("student"), handlers.UpdateStudentProfile(db))

	// --- Đối với Doanh nghiệp ---
	app.Get("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), handlers.GetBusinessProfile(db))
	app.Put("/api/profile/business", middleware.Protected(), middleware.RequireRole("business"), handlers.UpdateBusinessProfile(db))

	// Jobs & Applications Routes
	app.Post("/api/jobs", middleware.Protected(), middleware.RequireRole("business"), handlers.CreateJob(db))
	app.Get("/api/jobs", handlers.GetAvailableJobs(db)) // Công khai

	app.Post("/api/jobs/apply", middleware.Protected(), middleware.RequireRole("student"), handlers.ApplyJob(db))
	app.Put("/api/jobs/review-application", middleware.Protected(), middleware.RequireRole("business"), handlers.ReviewApplication(db))

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

	// Đặt cùng nhóm với các route của Student nhé Chanh
	app.Post("/api/applications/respond-offer",
		middleware.Protected(),
		middleware.RequireRole("student"),
		handlers.RespondToOffer(db),
	)

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

	app.Get("/api/wallet/me",
		middleware.Protected(),
		handlers.GetMyWallet(db),
	)
	// CHAT
	handlers.StartChatHub(db)

	// Middleware Upgrade WebSocket
	app.Use("/api/chat/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// REST API
	app.Get("/api/chat/history", handlers.GetChatHistory(db))

	// WebSocket
	app.Get("/api/chat/ws", websocket.New(handlers.HandleWS))

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
