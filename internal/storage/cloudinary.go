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
	uniquePublicID := fmt.Sprintf("%s/%d_%s", folder, time.Now().UnixNano(), baseName)

	resp, err := c.cld.Upload.Upload(context.Background(), bytes.NewReader(fileBytes), uploader.UploadParams{
		PublicID:     uniquePublicID,
		ResourceType: getCloudinaryResourceType(filename),
	})
	if err != nil {
		return "", fmt.Errorf("lỗi upload Cloudinary: %w", err)
	}

	log.Printf("🚀 [Cloudinary Engine]: Upload thành công lên Cloudinary: %s", resp.SecureURL)
	return resp.SecureURL, nil
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
