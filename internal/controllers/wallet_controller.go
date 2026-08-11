package controllers

import (
	"QuickWork/internal/services"

	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	walletService services.WalletService
}

func NewWalletController(walletService services.WalletService) *WalletController {
	return &WalletController{walletService: walletService}
}

// GetMyWallet GET /api/wallet/me
func (ctrl *WalletController) GetMyWallet(c *fiber.Ctx) error {
	userIDVal, ok := c.Locals("user_id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token không hợp lệ",
		})
	}

	userID := uint(userIDVal)
	wallet, transactions, err := ctrl.walletService.GetMyWallet(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Không thể tạo ví người dùng",
		})
	}

	return c.JSON(fiber.Map{
		"success":      true,
		"wallet":       wallet,
		"transactions": transactions,
	})
}
