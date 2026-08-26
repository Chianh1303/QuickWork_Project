package repositories

import (
	"QuickWork/internal/models"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetUserNotifications(userID uint, limit int) ([]models.Notification, error)
	GetUnreadCount(userID uint) (int64, error)
	MarkAsRead(id uint, userID uint) error
	MarkAllAsRead(userID uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetUserNotifications(userID uint, limit int) ([]models.Notification, error) {
	var notifs []models.Notification
	if limit <= 0 {
		limit = 20
	}
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Find(&notifs).Error
	return notifs, err
}

func (r *notificationRepository) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count).Error
	return count, err
}

func (r *notificationRepository) MarkAsRead(id uint, userID uint) error {
	return r.db.Model(&models.Notification{}).
		Where("id = ? AND user_id = ?", id, userID).
		Update("is_read", true).Error
}

func (r *notificationRepository) MarkAllAsRead(userID uint) error {
	return r.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true).Error
}
