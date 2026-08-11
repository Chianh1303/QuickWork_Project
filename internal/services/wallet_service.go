package services

import (
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"
)

type WalletService interface {
	GetMyWallet(userID uint) (*models.Wallet, []models.WalletTransaction, error)
}

type walletService struct {
	walletRepo repositories.WalletRepository
}

func NewWalletService(walletRepo repositories.WalletRepository) WalletService {
	return &walletService{walletRepo: walletRepo}
}

func (s *walletService) GetMyWallet(userID uint) (*models.Wallet, []models.WalletTransaction, error) {
	wallet, err := s.walletRepo.GetOrCreateWallet(userID)
	if err != nil {
		return nil, nil, err
	}
	transactions, err := s.walletRepo.GetTransactionsByWalletID(wallet.ID)
	if err != nil {
		return wallet, []models.WalletTransaction{}, nil
	}
	return wallet, transactions, nil
}
