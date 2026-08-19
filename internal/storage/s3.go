package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"QuickWork/internal/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageProvider interface {
	UploadFile(fileBytes []byte, filename string, folder string) (string, error)
	IsS3Active() bool
}

type s3StorageProvider struct {
	client     *s3.Client
	bucketName string
	region     string
	active     bool
}

func NewStorageProvider() StorageProvider {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := config.GetEnv("AWS_REGION", "ap-southeast-1")
	bucket := os.Getenv("AWS_S3_BUCKET")

	if accessKey == "" || secretKey == "" || bucket == "" {
		log.Println("⚠️ [AWS S3 Notice]: Chưa cấu hình AWS S3 credentials. Tự động chuyển sang Local Disk Storage.")
		return &s3StorageProvider{active: false}
	}

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		log.Printf("⚠️ [AWS S3 Notice]: Lỗi khởi tạo cấu hình AWS (%v). Tự động dùng Local Disk Storage.", err)
		return &s3StorageProvider{active: false}
	}

	client := s3.NewFromConfig(cfg)
	log.Printf("⚡ Kết nối AWS S3 Storage Provider thành công! (Bucket: %s)", bucket)

	return &s3StorageProvider{
		client:     client,
		bucketName: bucket,
		region:     region,
		active:     true,
	}
}

func (s *s3StorageProvider) IsS3Active() bool {
	return s.active
}

func (s *s3StorageProvider) UploadFile(fileBytes []byte, filename string, folder string) (string, error) {
	// FALLBACK: Nếu AWS S3 không khả dụng ➔ Lưu vào đĩa local
	if !s.active {
		return s.saveToLocalDisk(fileBytes, filename, folder)
	}

	// S3 UPLOAD: Upload trực tiếp lên đám mây AWS S3
	uniqueName := fmt.Sprintf("%s/%d_%s", folder, time.Now().UnixNano(), filename)
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &s.bucketName,
		Key:         &uniqueName,
		Body:        bytes.NewReader(fileBytes),
		ContentType: getContentType(filename),
	})
	if err != nil {
		log.Printf("⚠️ [AWS S3 Upload Notice]: Không thể upload lên AWS S3 (%v). Tự động Fallback sang Local Disk Storage...", err)
		return s.saveToLocalDisk(fileBytes, filename, folder)
	}

	s3URL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, uniqueName)
	log.Printf("🚀 [AWS S3 Engine]: Upload thành công lên đám mây S3: %s", s3URL)
	return s3URL, nil
}

func (s *s3StorageProvider) saveToLocalDisk(fileBytes []byte, filename string, folder string) (string, error) {
	targetDir := filepath.Join(config.UploadDir, folder)
	_ = os.MkdirAll(targetDir, os.ModePerm)

	uniqueName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)
	localPath := filepath.Join(targetDir, uniqueName)

	if err := os.WriteFile(localPath, fileBytes, 0644); err != nil {
		return "", fmt.Errorf("không thể lưu file đĩa local: %w", err)
	}

	publicURL := fmt.Sprintf("/uploads/%s/%s", folder, uniqueName)
	log.Printf("📁 [Local Storage]: Đã lưu file vào đĩa local: %s", publicURL)
	return publicURL, nil
}

func getContentType(filename string) *string {
	ext := filepath.Ext(filename)
	var contentType string
	switch ext {
	case ".pdf":
		contentType = "application/pdf"
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	default:
		contentType = "application/octet-stream"
	}
	return &contentType
}
