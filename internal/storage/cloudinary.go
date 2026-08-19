package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"QuickWork/internal/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type cloudinaryStorageProvider struct {
	cld    *cloudinary.Cloudinary
	active bool
}

func NewCloudinaryStorageProvider() *cloudinaryStorageProvider {
	cldURL := strings.TrimSpace(os.Getenv("CLOUDINARY_URL"))
	cloudName := strings.TrimSpace(os.Getenv("CLOUDINARY_CLOUD_NAME"))
	apiKey := strings.TrimSpace(os.Getenv("CLOUDINARY_API_KEY"))
	apiSecret := strings.TrimSpace(os.Getenv("CLOUDINARY_API_SECRET"))

	var cld *cloudinary.Cloudinary
	var err error

	if cldURL != "" {
		cld, err = cloudinary.NewFromURL(cldURL)
	} else if cloudName != "" && apiKey != "" && apiSecret != "" {
		cld, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	}

	if err != nil || cld == nil {
		return &cloudinaryStorageProvider{active: false}
	}

	log.Printf("⚡ Kết nối Cloudinary Storage Provider thành công! (Cloud: %s)", cld.Config.Cloud.CloudName)
	return &cloudinaryStorageProvider{
		cld:    cld,
		active: true,
	}
}

func (c *cloudinaryStorageProvider) IsActive() bool {
	return c.active
}

func (c *cloudinaryStorageProvider) UploadFile(fileBytes []byte, filename string, folder string) (string, error) {
	if !c.active || c.cld == nil {
		return "", fmt.Errorf("Cloudinary không khả dụng")
	}

	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	uniquePublicID := fmt.Sprintf("%d_%s", time.Now().UnixNano(), baseName)

	resp, err := c.cld.Upload.Upload(context.Background(), bytes.NewReader(fileBytes), uploader.UploadParams{
		Folder:       folder,
		PublicID:     uniquePublicID,
		ResourceType: getCloudinaryResourceType(filename),
	})
	if err != nil {
		log.Printf("⚠️ [Cloudinary Upload Notice]: Lỗi Cloudinary (%v). Tự động Fallback sang Local Disk Storage...", err)
		return c.saveToLocalDisk(fileBytes, filename, folder)
	}

	fileURL := resp.SecureURL
	if fileURL == "" {
		fileURL = resp.URL
	}
	if fileURL == "" && resp.PublicID != "" {
		fileURL = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/%s", c.cld.Config.Cloud.CloudName, resp.PublicID)
	}

	if fileURL == "" {
		log.Printf("⚠️ [Cloudinary Upload Notice]: Cloudinary URL rỗng (Resp: %+v). Tự động Fallback sang Local Disk Storage...", resp)
		return c.saveToLocalDisk(fileBytes, filename, folder)
	}

	log.Printf("🚀 [Cloudinary Engine]: Upload thành công lên Cloudinary: %s", fileURL)
	return fileURL, nil
}

func (c *cloudinaryStorageProvider) saveToLocalDisk(fileBytes []byte, filename string, folder string) (string, error) {
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

func getCloudinaryResourceType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".pdf":
		return "raw"
	default:
		return "auto"
	}
}
