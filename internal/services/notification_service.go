package services

import (
	"QuickWork/internal/models"
	"QuickWork/internal/queue"
	"QuickWork/internal/repositories"
)

type NotificationService interface {
	CreateNotification(userID uint, title string, message string, notifType string, refID uint) error
	GetUserNotifications(userID uint, limit int) ([]models.Notification, int64, error)
	MarkAsRead(id uint, userID uint) error
	MarkAllAsRead(userID uint) error
}

type notificationService struct {
	repo      repositories.NotificationRepository
	rmqClient queue.RabbitMQClient
}

func NewNotificationService(repo repositories.NotificationRepository, rmqClient ...queue.RabbitMQClient) NotificationService {
	var rmq queue.RabbitMQClient
	if len(rmqClient) > 0 {
		rmq = rmqClient[0]
	}
	return &notificationService{
		repo:      repo,
		rmqClient: rmq,
	}
}

func (s *notificationService) CreateNotification(userID uint, title string, message string, notifType string, refID uint) error {
	notif := &models.Notification{
		UserID:      userID,
		Title:       title,
		Message:     message,
		Type:        notifType,
		ReferenceID: refID,
		IsRead:      false,
	}

	err := s.repo.CreateNotification(notif)
	if err != nil {
		return err
	}

	// Optionally publish to RabbitMQ Notification queue
	if s.rmqClient != nil && s.rmqClient.IsAvailable() {
		_ = s.rmqClient.Publish(queue.QueueNotifications, queue.NotificationPayload{
			UserID:  userID,
			Title:   title,
			Message: message,
		})
	}

	return nil
}

func (s *notificationService) GetUserNotifications(userID uint, limit int) ([]models.Notification, int64, error) {
	notifs, err := s.repo.GetUserNotifications(userID, limit)
	if err != nil {
		return nil, 0, err
	}

	unreadCount, err := s.repo.GetUnreadCount(userID)
	if err != nil {
		unreadCount = 0
	}

	return notifs, unreadCount, nil
}

func (s *notificationService) MarkAsRead(id uint, userID uint) error {
	return s.repo.MarkAsRead(id, userID)
}

func (s *notificationService) MarkAllAsRead(userID uint) error {
	return s.repo.MarkAllAsRead(userID)
}
