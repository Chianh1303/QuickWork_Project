package repositories

import (
	"QuickWork/internal/models"
	"gorm.io/gorm"
)

type ChatRepository interface {
	GetMessagesByApplicationID(appID uint) ([]models.Message, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) GetMessagesByApplicationID(appID uint) ([]models.Message, error) {
	var messages []models.Message
	if err := r.db.
		Where("application_id = ?", appID).
		Order("created_at asc").
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
