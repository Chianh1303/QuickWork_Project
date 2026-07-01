package config

const (
	MySQLDSN    = "root:root@123@tcp(127.0.0.1:3306)/quickwork_db?charset=utf8mb4&parseTime=True&loc=Local"
	ServerAddr  = ":3000"
	UploadDir   = "./uploads"
	AvatarDir   = "./uploads/avatars"
	CVDir       = "./uploads/cvs"
	CORSOrigins = "http://localhost:3001,http://127.0.0.1:3001"
)
