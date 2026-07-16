package handlers

import (
	"math"
	"strconv"
	"strings"

	"QuickWork/internal/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
				Where("users.status = ?", "pending")

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
				users.created_at
			`).
			Order("users.created_at DESC").
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
