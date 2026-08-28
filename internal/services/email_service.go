package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"

	"QuickWork/internal/config"
)

type EmailService interface {
	SendOTPEmail(toEmail string, otpCode string) error
}

type emailService struct {
	smtpHost     string
	smtpPort     string
	smtpSecure   bool
	authEmail    string
	authPassword string
	brevoAPIKey  string
	resendAPIKey string
}

func NewEmailService() EmailService {
	port := config.GetEnv("SMTP_PORT", "587")
	secureEnv := strings.ToLower(config.GetEnv("SMTP_SECURE", config.GetEnv("SMTP_SSL", "false")))
	isSecure := secureEnv == "true" || secureEnv == "1" || port == "465"

	return &emailService{
		smtpHost:     config.GetEnv("SMTP_HOST", "smtp.gmail.com"),
		smtpPort:     port,
		smtpSecure:   isSecure,
		authEmail:    strings.TrimSpace(os.Getenv("SMTP_EMAIL")),
		authPassword: strings.TrimSpace(os.Getenv("SMTP_PASSWORD")),
		brevoAPIKey:  strings.TrimSpace(os.Getenv("BREVO_API_KEY")),
		resendAPIKey: strings.TrimSpace(os.Getenv("RESEND_API_KEY")),
	}
}

func (e *emailService) SendOTPEmail(toEmail string, otpCode string) error {
	subject := "Ma xac thuc khoi phuc mat khau - QuickWork"
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

	// =========================================================================
	// OPTION 1: HTTP REST API via Brevo (Port 443 - Cannot be blocked by Render)
	// =========================================================================
	if e.brevoAPIKey != "" {
		log.Printf("🚀 [Email API]: Đang gửi mã OTP qua Brevo HTTP REST API (Port 443) đến %s...", toEmail)
		err := sendViaBrevo(e.brevoAPIKey, e.authEmail, toEmail, subject, body)
		if err == nil {
			log.Printf("🎉 [Email API]: Gửi mã OTP thành công 100%% qua Brevo HTTP REST API đến Gmail: %s", toEmail)
			return nil
		}
		log.Printf("⚠️ [Email API Error]: Brevo HTTP API lỗi (%v). Thử nghiệm luồng kế tiếp...", err)
	}

	// OPTION 1B: HTTP REST API via Resend (Port 443)
	if e.resendAPIKey != "" {
		log.Printf("🚀 [Email API]: Đang gửi mã OTP qua Resend HTTP REST API (Port 443) đến %s...", toEmail)
		err := sendViaResend(e.resendAPIKey, e.authEmail, toEmail, subject, body)
		if err == nil {
			log.Printf("🎉 [Email API]: Gửi mã OTP thành công 100%% qua Resend HTTP REST API đến Gmail: %s", toEmail)
			return nil
		}
		log.Printf("⚠️ [Email API Error]: Resend HTTP API lỗi (%v). Thử nghiệm luồng kế tiếp...", err)
	}

	// =========================================================================
	// OPTION 2: SMTP Connection with 30-second Timeout & Detailed TLS Tracing
	// =========================================================================
	if e.authEmail == "" || e.authPassword == "" {
		log.Printf("❌ [Email Error]: CHƯA CẤU HÌNH SMTP_EMAIL / SMTP_PASSWORD TRÊN RENDER ENVIRONMENT! [Mã OTP Console Fallback]: %s -> %s", toEmail, otpCode)
		return fmt.Errorf("chưa cấu hình SMTP_EMAIL / SMTP_PASSWORD trên Render Environment")
	}

	log.Printf("📧 [SMTP Engine]: Đang gửi mã OTP đến %s (Host: %s:%s | 30s Timeout)...", toEmail, e.smtpHost, e.smtpPort)

	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("QuickWork Platform <%s>", e.authEmail)
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	msgBytes := []byte(message)

	// Mode 2A: Port 587 STARTTLS (30s Timeout)
	if !e.smtpSecure && e.smtpPort == "587" {
		err587 := sendMailPort587(e.smtpHost, "587", e.authEmail, e.authPassword, toEmail, msgBytes, 30*time.Second)
		if err587 == nil {
			log.Printf("🎉 [SMTP Engine]: Gửi mã OTP thành công 100%% qua Port 587 STARTTLS đến Gmail: %s", toEmail)
			return nil
		}
		log.Printf("⚠️ [SMTP Engine Notice]: Cổng 587 STARTTLS báo lỗi bắt tay/timeout (%v). Chuyển sang Port 465 SSL...", err587)
	}

	// Mode 2B: Port 465 Direct SSL (30s Timeout)
	err465 := sendMailPort465(e.smtpHost, "465", e.authEmail, e.authPassword, toEmail, msgBytes, 30*time.Second)
	if err465 == nil {
		log.Printf("🎉 [SMTP Engine]: Gửi mã OTP thành công 100%% qua Port 465 SSL đến Gmail: %s", toEmail)
		return nil
	}

	log.Printf("❌ [SMTP Error Final]: Cả Port 587 và 465 đều thất bại (%v). [Mã OTP Fallback]: %s -> %s", err465, toEmail, otpCode)
	return err465
}

// Brevo HTTP REST API Handler (Port 443)
func sendViaBrevo(apiKey, fromEmail, toEmail, subject, htmlBody string) error {
	if fromEmail == "" {
		fromEmail = "chianhn567@gmail.com"
	}
	url := "https://api.brevo.com/v3/smtp/email"
	payload := map[string]interface{}{
		"sender": map[string]string{
			"name":  "QuickWork Platform",
			"email": fromEmail,
		},
		"to": []map[string]string{
			{"email": toEmail},
		},
		"subject":     subject,
		"htmlContent": htmlBody,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", apiKey)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	bodyResp, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("Brevo HTTP API Error (Status %d): %s", resp.StatusCode, string(bodyResp))
}

// Resend HTTP REST API Handler (Port 443)
func sendViaResend(apiKey, fromEmail, toEmail, subject, htmlBody string) error {
	url := "https://api.resend.com/emails"
	payload := map[string]interface{}{
		"from":    "QuickWork <onboarding@resend.dev>",
		"to":      []string{toEmail},
		"subject": subject,
		"html":    htmlBody,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	bodyResp, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("Resend HTTP API Error (Status %d): %s", resp.StatusCode, string(bodyResp))
}

// Send Mail via Port 587 (STARTTLS) with Custom Timeout
func sendMailPort587(host, port, from, password, to string, body []byte, timeout time.Duration) error {
	addr := host + ":" + port
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return fmt.Errorf("lỗi kết nối TCP port 587 (%w)", err)
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("lỗi khởi tạo SMTP client 587 (%w)", err)
	}
	defer c.Quit()

	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName: host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			return fmt.Errorf("lỗi bắt tay TLS StartTLS (%w)", err)
		}
	}

	auth := smtp.PlainAuth("", from, password, host)
	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("lỗi xác thực mật khẩu Gmail (%w)", err)
	}

	if err = c.Mail(from); err != nil {
		return fmt.Errorf("lỗi khai báo mail From (%w)", err)
	}
	if err = c.Rcpt(to); err != nil {
		return fmt.Errorf("lỗi khai báo mail To (%w)", err)
	}

	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("lỗi mở stream Data (%w)", err)
	}
	if _, err = w.Write(body); err != nil {
		return fmt.Errorf("lỗi ghi nội dung email (%w)", err)
	}
	return w.Close()
}

// Send Mail via Port 465 (Direct SSL) with Custom Timeout
func sendMailPort465(host, port, from, password, to string, body []byte, timeout time.Duration) error {
	dialer := &net.Dialer{Timeout: timeout}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", host+":"+port, tlsConfig)
	if err != nil {
		return fmt.Errorf("lỗi kết nối SSL port 465 (%w)", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("lỗi khởi tạo Client SSL 465 (%w)", err)
	}
	defer client.Quit()

	auth := smtp.PlainAuth("", from, password, host)
	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("lỗi xác thực mật khẩu SSL (%w)", err)
	}

	if err := client.Mail(from); err != nil {
		return err
	}
	if err := client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}
	if _, err := w.Write(body); err != nil {
		return err
	}
	return w.Close()
}
