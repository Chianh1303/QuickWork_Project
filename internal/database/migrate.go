package database

import (
	"QuickWork/internal/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Business{},
		&models.Job{},
		&models.Application{},
		&models.Message{},
		&models.Attendance{},
		&models.Wallet{},
		&models.WalletTransaction{},
		&models.Review{},
	)
}
