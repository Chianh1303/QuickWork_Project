package handlers

import (
	"QuickWork/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetMyWallet(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDVal, ok := c.Locals("user_id").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token không hợp lệ",
			})
		}

		userID := uint(userIDVal)

		var wallet models.Wallet
		if err := db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
			wallet = models.Wallet{
				UserID:  userID,
				Balance: 0,
			}

			if err := db.Create(&wallet).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Không thể tạo ví người dùng",
				})
			}
		}

		var transactions []models.WalletTransaction
		db.Where("wallet_id = ?", wallet.ID).
			Order("created_at DESC").
			Find(&transactions)

		return c.JSON(fiber.Map{
			"success":      true,
			"wallet":       wallet,
			"transactions": transactions,
		})
	}
}
