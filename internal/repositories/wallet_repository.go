package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	GetOrCreateWallet(userID uint) (*models.Wallet, error)
	GetTransactionsByWalletID(walletID uint) ([]models.WalletTransaction, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) GetOrCreateWallet(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		wallet = models.Wallet{
			UserID:  userID,
			Balance: 0,
		}
		if err := r.db.Create(&wallet).Error; err != nil {
			return nil, err
		}
	}
	return &wallet, nil
}

func (r *walletRepository) GetTransactionsByWalletID(walletID uint) ([]models.WalletTransaction, error) {
	var transactions []models.WalletTransaction
	r.db.Where("wallet_id = ?", walletID).
		Order("created_at DESC").
		Find(&transactions)
	return transactions, nil
}
