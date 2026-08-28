package services

import (
	"crypto/tls"
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

	log.Printf("📧 [Email Engine]: Đang khởi chạy tiến trình gửi mã OTP đến %s (User: %s)...", toEmail, e.authEmail)

	body := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<style>
		body { font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; background-color: #0f172a; margin: 0; padding: 20px; color: #f8fafc; }
		.card { max-width: 500px; margin: 0 auto; background: #1e293b; border: 1px solid rgba(34, 211, 238, 0.2); border-radius: 16px; padding: 32px; box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.5); }
		.logo { font-size: 24px; font-weight: 900; color: #22d3ee; margin-bottom: 24px; text-align: center; letter-spacing: -0.5px; }
		.title { font-size: 18px; font-weight: 800; color: #ffffff; text-align: center; margin-bottom: 8px; }
		.desc { font-size: 13px; color: #94a3b8; text-align: center; margin-bottom: 24px; line-height: 1.5; }
		.otp-box { background: rgba(34, 211, 238, 0.1); border: 2px dashed #22d3ee; border-radius: 12px; font-size: 36px; font-weight: 900; color: #38bdf8; letter-spacing: 8px; text-align: center; padding: 18px; margin: 24px 0; }
		.warning { font-size: 12px; color: #f43f5e; text-align: center; font-weight: 600; margin-top: 16px; }
		.footer { font-size: 11px; color: #64748b; text-align: center; margin-top: 32px; border-t: 1px solid rgba(255,255,255,0.05); padding-top: 16px; }
	</style>
</head>
<body>
	<div class="card">
		<div class="logo">⚡ QUICKWORK PLATFORM</div>
		<div class="title">Mã Xác Thực Khôi Phục Mật Khẩu</div>
		<div class="desc">Bạn đã gửi yêu cầu đặt lại mật khẩu cho tài khoản <strong style="color: #38bdf8;">%s</strong>. Vui lòng nhập mã OTP bên dưới để hoàn tất:</div>
		<div class="otp-box">%s</div>
		<div class="warning">⚠️ Mã xác thực có hiệu lực trong 5 phút. Tuyệt đối không chia sẻ mã này cho người khác.</div>
		<div class="footer">Nếu bạn không yêu cầu mã này, vui lòng bỏ qua email này.<br>&copy; 2026 QuickWork Platform. All rights reserved.</div>
	</div>
</body>
</html>`, toEmail, otpCode)

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("QuickWork Platform <%s>", e.authEmail)
	headers["To"] = toEmail
	headers["Subject"] = "Ma xac thuc khoi phuc mat khau - QuickWork"
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	msgBytes := []byte(message)
	auth := smtp.PlainAuth("", e.authEmail, e.authPassword, e.smtpHost)

	// Mode 1: Try Port 587 Standard SendMail
	err587 := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.authEmail, []string{toEmail}, msgBytes)
	if err587 == nil {
		log.Printf("🎉 [Email Engine]: Gửi mã OTP thành công 100%% qua Port %s đến Gmail: %s", e.smtpPort, toEmail)
		return nil
	}

	log.Printf("⚠️ [Email Engine]: Gửi qua Port %s không thành công (%v). Chuyển sang kết nối Port 465 SSL/TLS...", e.smtpPort, err587)

	// Mode 2: Try Port 465 Direct TLS SSL Connection
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.smtpHost,
	}

	conn, errTls := tls.Dial("tcp", e.smtpHost+":465", tlsconfig)
	if errTls != nil {
		log.Printf("❌ [Email Engine]: Lỗi kết nối Port 465 TLS (%v). [Mã OTP Fallback]: %s -> %s", errTls, toEmail, otpCode)
		return errTls
	}
	defer conn.Close()

	client, errClient := smtp.NewClient(conn, e.smtpHost)
	if errClient != nil {
		log.Printf("❌ [Email Engine]: Lỗi tạo Client SMTP Port 465 (%v). [Mã OTP Fallback]: %s -> %s", errClient, toEmail, otpCode)
		return errClient
	}
	defer client.Quit()

	if errAuth := client.Auth(auth); errAuth != nil {
		log.Printf("❌ [Email Engine]: Lỗi Mật khẩu App Gmail / Auth (%v). [Mã OTP Fallback]: %s -> %s", errAuth, toEmail, otpCode)
		return errAuth
	}

	if errMail := client.Mail(e.authEmail); errMail != nil {
		return errMail
	}
	if errRcpt := client.Rcpt(toEmail); errRcpt != nil {
		return errRcpt
	}

	w, errData := client.Data()
	if errData != nil {
		return errData
	}
	_, errWrite := w.Write(msgBytes)
	if errWrite != nil {
		return errWrite
	}
	_ = w.Close()

	log.Printf("🎉 [Email Engine]: Gửi mã OTP thành công 100%% qua Port 465 SSL đến Gmail: %s", toEmail)
	return nil
}
