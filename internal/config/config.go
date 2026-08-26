package config

import (
	"fmt"
	"os"
)

var (
	MySQLDSN    = getMySQLDSN()
	ServerAddr  = GetEnv("SERVER_ADDR", ":3000")
	UploadDir   = "./uploads"
	AvatarDir   = "./uploads/avatars"
	CVDir       = "./uploads/cvs"
	CORSOrigins = GetEnv("CORS_ORIGINS", "http://localhost:3001,http://127.0.0.1:3001")
	RabbitMQURL = GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
)

func getMySQLDSN() string {
	dbUser := GetEnv("DB_USER", "root")
	dbPass := GetEnv("DB_PASSWORD", "root@123")
	dbHost := GetEnv("DB_HOST", "127.0.0.1")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "quickwork_db")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
}

func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
