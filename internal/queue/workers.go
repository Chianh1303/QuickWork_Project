package queue

import (
	"encoding/json"
	"log"
)

const (
	QueueCVParsing     = "cv_parsing_queue"
	QueueNotifications = "notification_queue"
	QueueEmailOTP      = "email_otp_queue"
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

type EmailOTPPayload struct {
	Email   string `json:"email"`
	OTPCode string `json:"otp_code"`
}

type CVParsingHandler func(userID uint, cvURL string) error
type EmailOTPHandler func(email string, otpCode string) error

func RegisterWorkers(rmq RabbitMQClient, cvHandler CVParsingHandler, otpHandler EmailOTPHandler) {
	if rmq == nil || !rmq.IsAvailable() {
		return
	}

	// 1. Worker xử lý giải mã & phân tích CV ngầm thực tế
	_ = rmq.Consume(QueueCVParsing, func(body []byte) error {
		var payload CVParsingPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		log.Printf("⚙️ [Background Worker]: Bắt đầu giải mã CV PDF ngầm cho User ID %d (URL: %s)...", payload.UserID, payload.CvURL)
		if cvHandler != nil {
			if err := cvHandler(payload.UserID, payload.CvURL); err != nil {
				log.Printf("❌ [Background Worker Error]: %v", err)
				return err
			}
		}
		log.Printf("✅ [Background Worker]: Đã hoàn thành phân tích AI CV ngầm cho User ID %d!", payload.UserID)
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

	// 3. Worker xử lý gửi Email OTP ngầm qua RabbitMQ
	_ = rmq.Consume(QueueEmailOTP, func(body []byte) error {
		var payload EmailOTPPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		log.Printf("📧 [RabbitMQ Email Worker]: Bắt đầu gửi Email OTP ngầm đến %s...", payload.Email)
		if otpHandler != nil {
			if err := otpHandler(payload.Email, payload.OTPCode); err != nil {
				log.Printf("❌ [RabbitMQ Email Error]: %v", err)
				return err
			}
		}
		log.Printf("✅ [RabbitMQ Email Worker]: Đã gửi xong Email OTP ngầm cho %s!", payload.Email)
		return nil
	})
}
