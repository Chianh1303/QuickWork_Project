package config

import (
	"fmt"
	"os"
	"strings"
)

var (
	MySQLDSN    = getMySQLDSN()
	ServerAddr  = getServerAddr()
	UploadDir   = "./uploads"
	AvatarDir   = "./uploads/avatars"
	CVDir       = "./uploads/cvs"
	CORSOrigins = GetEnv("CORS_ORIGINS", "*")
	RabbitMQURL = GetEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
)

func getServerAddr() string {
	port := os.Getenv("SERVER_ADDR")
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		port = "8080"
	}
	port = strings.TrimPrefix(port, ":")
	return "0.0.0.0:" + port
}

func getMySQLDSN() string {
	dbUser := GetEnv("DB_USER", "root")
	dbPass := GetEnv("DB_PASSWORD", "root@123")
	dbHost := GetEnv("DB_HOST", "127.0.0.1")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "quickwork_db")
	dbParams := GetEnv("DB_PARAMS", "charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPass, dbHost, dbPort, dbName, dbParams)
}

func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
