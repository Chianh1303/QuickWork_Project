package handlers

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"QuickWork/internal/dto"
	"QuickWork/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAdminDashboardStats(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var stats dto.AdminDashboardStats

		if err := db.Model(&models.User{}).Where("role = ?", "student").Count(&stats.TotalStudents).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}
		if err := db.Model(&models.User{}).Where("role = ?", "business").Count(&stats.TotalBusinesses).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}
		if err := db.Model(&models.User{}).Where("role = ? AND status = ?", "business", "pending").Count(&stats.PendingBusinesses).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}
		if err := db.Model(&models.Job{}).Count(&stats.TotalJobs).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}
		if err := db.Model(&models.Job{}).Where("status = ?", "pending").Count(&stats.PendingJobs).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}

		// total_disbursed là tổng tiền lương đã giải ngân qua hệ thống, không phải doanh thu/lợi nhuận QuickWork.
		if err := db.Model(&models.WalletTransaction{}).
			Where("type = ?", "salary").
			Select("COALESCE(SUM(amount), 0)").
			Scan(&stats.TotalDisbursed).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get dashboard stats"})
		}

		return c.JSON(stats)
	}
}

func GetPendingBusinesses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, err := strconv.Atoi(c.Query("page", "1"))
		if err != nil || page < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "page must be greater than or equal to 1",
			})
		}

		limit, err := strconv.Atoi(c.Query("limit", "10"))
		if err != nil || limit < 1 || limit > 100 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "limit must be between 1 and 100",
			})
		}

		search := strings.TrimSpace(c.Query("search"))
		offset := (page - 1) * limit

		buildQuery := func() *gorm.DB {
			query := db.
				Table("businesses").
				Joins("JOIN users ON users.id = businesses.user_id").
				Where("users.role = ? AND users.status = ?", "business", "pending")

			if search != "" {
				keyword := "%" + search + "%"
				query = query.Where(
					"(businesses.company_name LIKE ? OR businesses.tax_code LIKE ? OR users.email LIKE ?)",
					keyword,
					keyword,
					keyword,
				)
			}

			return query
		}

		var total int64
		if err := buildQuery().Count(&total).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to count pending businesses",
			})
		}

		items := make([]dto.PendingBusinessItem, 0)
		if err := buildQuery().
			Select(`
				businesses.id AS business_id,
				businesses.company_name,
				businesses.tax_code,
				users.email,
				users.status,
				businesses.created_at
			`).
			Order("businesses.created_at DESC").
			Limit(limit).
			Offset(offset).
			Scan(&items).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to get pending businesses",
			})
		}

		totalPages := int64(0)
		if total > 0 {
			totalPages = int64(math.Ceil(float64(total) / float64(limit)))
		}

		return c.JSON(dto.PendingBusinessListResponse{
			Items: items,
			Pagination: dto.PaginationMeta{
				Page:       page,
				Limit:      limit,
				Total:      total,
				TotalPages: totalPages,
			},
		})
	}
}

func GetBusinessKYBDetail(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		businessID, err := strconv.Atoi(c.Params("id"))
		if err != nil || businessID < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "business id is invalid"})
		}

		var detail dto.BusinessKYBDetail
		err = db.
			Table("businesses").
			Joins("JOIN users ON users.id = businesses.user_id").
			Where("businesses.id = ? AND users.role = ?", businessID, "business").
			Select(`
				businesses.id AS business_id,
				businesses.user_id,
				businesses.company_name,
				businesses.tax_code,
				users.email,
				businesses.phone,
				businesses.address,
				businesses.logo_url,
				users.status,
				businesses.is_verified,
				businesses.reject_reason,
				businesses.created_at,
				businesses.reviewed_at
			`).
			Scan(&detail).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to get business detail"})
		}
		if detail.BusinessID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "business not found"})
		}

		return c.JSON(detail)
	}
}

func ReviewBusinessKYB(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		businessID, err := strconv.Atoi(c.Params("id"))
		if err != nil || businessID < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "business id is invalid"})
		}

		adminID, ok := localUserID(c)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token không hợp lệ"})
		}

		var req dto.ReviewBusinessRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Dữ liệu không hợp lệ"})
		}

		req.Decision = strings.TrimSpace(req.Decision)
		req.RejectReason = strings.TrimSpace(req.RejectReason)
		if req.Decision != "approved" && req.Decision != "rejected" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "decision chỉ được là approved hoặc rejected"})
		}
		if req.Decision == "rejected" && len([]rune(req.RejectReason)) < 10 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Lý do từ chối phải có ít nhất 10 ký tự"})
		}

		now := time.Now()

		// Transaction bảo đảm users.status và businesses.is_verified luôn cập nhật cùng nhau hoặc rollback cùng nhau.
		err = db.Transaction(func(tx *gorm.DB) error {
			var business models.Business
			// Row lock ngăn hai Admin ghi đè quyết định KYB của nhau trên cùng một hồ sơ.
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&business, uint(businessID)).Error; err != nil {
				return err
			}

			var user models.User
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("id = ? AND role = ?", business.UserID, "business").
				First(&user).Error; err != nil {
				return err
			}

			if user.Status != "pending" {
				return fiber.NewError(fiber.StatusConflict, "Hồ sơ doanh nghiệp đã được xử lý")
			}

			nextStatus := "approved"
			isVerified := true
			rejectReason := ""
			if req.Decision == "approved" {
				nextStatus = "approved"
				isVerified = true
				rejectReason = ""
			} else {
				nextStatus = "rejected"
				isVerified = false
				rejectReason = req.RejectReason
			}

			if err := tx.Model(&user).Update("status", nextStatus).Error; err != nil {
				return err
			}

			// Chỉ update các field KYB cần thiết để không ghi đè created_at cũ có thể đang là zero date.
			if err := tx.Model(&business).Updates(map[string]interface{}{
				"is_verified":   isVerified,
				"reject_reason": rejectReason,
				"reviewed_by":   adminID,
				"reviewed_at":   now,
			}).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				return c.Status(fiberErr.Code).JSON(fiber.Map{"message": fiberErr.Message})
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "business not found"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to review business"})
		}

		return c.JSON(fiber.Map{
			"message":  "Đã cập nhật kết quả KYB doanh nghiệp",
			"decision": req.Decision,
		})
	}
}

func localUserID(c *fiber.Ctx) (uint, bool) {
	switch value := c.Locals("user_id").(type) {
	case float64:
		return uint(value), true
	case uint:
		return value, true
	case int:
		return uint(value), true
	default:
		return 0, false
	}
}
