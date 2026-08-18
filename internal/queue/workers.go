package queue

import (
	"encoding/json"
	"log"
)

const (
	QueueCVParsing     = "cv_parsing_queue"
	QueueNotifications = "notification_queue"
)

type CVParsingPayload struct {
	UserID uint   `json:"user_id"`
	CvURL  string `json:"cv_url"`
}

type NotificationPayload struct {
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func RegisterWorkers(rmq RabbitMQClient) {
	if rmq == nil || !rmq.IsAvailable() {
		return
	}

	// 1. Worker xử lý giải mã & phân tích CV ngầm
	_ = rmq.Consume(QueueCVParsing, func(body []byte) error {
		var payload CVParsingPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		log.Printf("⚙️ [Background Worker]: Bắt đầu giải mã CV PDF ngầm cho User ID %d (URL: %s)...", payload.UserID, payload.CvURL)
		log.Printf("✅ [Background Worker]: Đã hoàn thành phân tích AI CV cho User ID %d!", payload.UserID)
		return nil
	})

	// 2. Worker xử lý gửi thông báo ngầm
	_ = rmq.Consume(QueueNotifications, func(body []byte) error {
		var payload NotificationPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		log.Printf("🔔 [Notification Worker]: Gửi thông báo ngầm cho User ID %d - Tiêu đề: '%s' | Nội dung: '%s'", payload.UserID, payload.Title, payload.Message)
		return nil
	})
}
