package config

import "os"

const (
	MySQLDSN    = "root:root@123@tcp(127.0.0.1:3306)/quickwork_db?charset=utf8mb4&parseTime=True&loc=Local"
	ServerAddr  = ":3000"
	UploadDir   = "./uploads"
	AvatarDir   = "./uploads/avatars"
	CVDir       = "./uploads/cvs"
	CORSOrigins = "http://localhost:3001,http://127.0.0.1:3001"
	RabbitMQURL = "amqp://guest:guest@localhost:5672/"
)

func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
