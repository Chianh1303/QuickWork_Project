package models

// 4. Bảng Tin tuyển dụng (Khớp chuẩn chỉnh với DBML của bạn)
type Job struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	BusinessID  uint     `gorm:"not null" json:"business_id"` // Khóa ngoại nối sang Business
	Business    Business `gorm:"foreignKey:BusinessID" json:"business"`
	Title       string   `gorm:"type:varchar(255);not null" json:"title"`
	Description string   `gorm:"type:text;not null" json:"description"`
	Location    string   `gorm:"type:varchar(255);not null" json:"location"`
	Salary      float64  `gorm:"type:decimal(15,2);not null" json:"salary"`
	Slots       int      `gorm:"not null;default:1" json:"slots"`
	Status      string   `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, closed
	WorkingDate string   `gorm:"type:varchar(100)" json:"working_date"`
	Category    string   `json:"category"` // Ngành nghề (e.g., "IT", "Marketing")
	JobType     string   `json:"job_type"`
}
