package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"QuickWork/internal/config"
)

type StorageProvider interface {
	UploadFile(fileBytes []byte, filename string, folder string) (string, error)
}

type localStorageProvider struct{}

func NewStorageProvider() StorageProvider {
	// 1. Ưu tiên kiểm tra kết nối Cloudinary Storage
	cldProvider := NewCloudinaryStorageProvider()
	if cldProvider.IsActive() {
		return cldProvider
	}

	// 2. Fallback: Dùng Local Disk Storage
	log.Println("⚠️ [Storage Notice]: Chưa cấu hình Cloudinary credentials. Tự động dùng Local Disk Storage.")
	return &localStorageProvider{}
}

func (l *localStorageProvider) UploadFile(fileBytes []byte, filename string, folder string) (string, error) {
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
