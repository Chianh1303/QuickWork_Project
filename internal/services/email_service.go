package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"QuickWork/internal/config"
)

type EmailService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}

type emailService struct {
	smtpHost     string
	smtpPort     string
	authEmail    string
	authPassword string
}

func NewEmailService() EmailService {
	return &emailService{
		smtpHost:     config.GetEnv("SMTP_HOST", "smtp.gmail.com"),
		smtpPort:     config.GetEnv("SMTP_PORT", "587"),
		authEmail:    strings.TrimSpace(os.Getenv("SMTP_EMAIL")),
		authPassword: strings.TrimSpace(os.Getenv("SMTP_PASSWORD")),
	}
}

func (e *emailService) SendOTPEmail(toEmail string, otpCode string) error {
	if e.authEmail == "" || e.authPassword == "" {
		log.Printf("⚠️ [Email Notice]: Chưa cấu hình SMTP_EMAIL / SMTP_PASSWORD trong .env. [Mã OTP Console Fallback]: %s -> %s", toEmail, otpCode)
		return nil
	}

	subject := "Subject: [QuickWork] Ma xac nhan khoi phuc mat khau\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; background-color: #f4f6f8; margin: 0; padding: 20px; }
			.card { max-width: 500px; margin: 0 auto; background: #ffffff; border-radius: 12px; padding: 30px; box-shadow: 0 4px 15px rgba(0,0,0,0.05); }
			.logo { font-size: 24px; font-weight: bold; color: #4F46E5; margin-bottom: 20px; text-align: center; }
			.otp-box { background: #EEF2FF; border: 2px dashed #6366F1; border-radius: 8px; font-size: 32px; font-weight: bold; color: #4F46E5; letter-spacing: 6px; text-align: center; padding: 15px; margin: 20px 0; }
			.footer { font-size: 12px; color: #6B7280; text-align: center; margin-top: 25px; }
		</style>
	</head>
	<body>
		<div class="card">
			<div class="logo">⚡ QuickWork Platform</div>
			<h2>Mã Xác Thực Khôi Phục Mật Khẩu</h2>
			<p>Bạn đã yêu cầu đặt lại mật khẩu cho tài khoản <strong>%s</strong>.</p>
			<p>Vui lòng nhập mã OTP gồm 6 chữ số dưới đây để tiếp tục:</p>
			<div class="otp-box">%s</div>
			<p>⚠️ Mã này có hiệu lực trong vòng <strong>5 phút</strong>. Vui lòng không chia sẻ mã này cho bất kỳ ai.</p>
			<div class="footer">Nếu bạn không yêu cầu mã này, vui lòng bỏ qua email này.<br>&copy; 2026 QuickWork Platform. All rights reserved.</div>
		</div>
	</body>
	</html>
	`, toEmail, otpCode)

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", e.authEmail, e.authPassword, e.smtpHost)

	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.authEmail, []string{toEmail}, msg)
	if err != nil {
		log.Printf("⚠️ [Email Error]: Lỗi gửi mail qua SMTP (%v). [Mã OTP Fallback]: %s -> %s", err, toEmail, otpCode)
		return err
	}

	log.Printf("📧 [Email Engine]: Đã gửi mã OTP thành công đến Gmail: %s", toEmail)
	return nil
}
