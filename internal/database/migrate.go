package database

import (
	"QuickWork/internal/models"
	"fmt"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Kiểm tra dữ liệu trùng lặp (Duplicate Detection) trên bảng tickets nếu bảng đã tồn tại
	if db.Migrator().HasTable(&models.Ticket{}) {
		var duplicateCount int64
		err := db.Raw(`
			SELECT COUNT(*) FROM (
				SELECT application_id, reporter_id 
				FROM tickets 
				GROUP BY application_id, reporter_id 
				HAVING COUNT(*) > 1
			) AS duplicates
		`).Scan(&duplicateCount).Error

		if err == nil && duplicateCount > 0 {
			return fmt.Errorf(
				"❌ MIGRATION ABORTED: Phát hiện %d nhóm Ticket bị trùng lặp (application_id, reporter_id) trong CSDL! Vui lòng kiểm tra và dọn dẹp thủ công trước khi tạo UNIQUE INDEX",
				duplicateCount,
			)
		}
	}

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
		&models.Ticket{},
		&models.Category{},
		&models.Skill{},
		&models.CVEvaluation{},
		&models.Notification{},
		&models.SavedJob{},
	)
}
