package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password  string     `gorm:"not null;type:varchar(255)" json:"-"`
	Role      string     `gorm:"type:varchar(20);not null" json:"role"` // admin, business, student
	Status    string     `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Balance   float64    `gorm:"type:decimal(15,2);default:0.00" json:"balance"` // Khớp với decimal của bạn
	CreatedAt time.Time  `json:"created_at"`
	
	Student   *Student   `json:"student,omitempty"`
	Business  *Business  `json:"business,omitempty"`
}
type Student struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `gorm:"not null;unique" json:"user_id"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`

	FullName string `gorm:"type:varchar(100);not null" json:"full_name"`

	Phone string `gorm:"type:varchar(20)" json:"phone"`

	Gender string `gorm:"type:varchar(10)" json:"gender"`

	AvatarUrl string `gorm:"type:varchar(255)" json:"avatar_url"`

	Skills string `gorm:"type:text" json:"skills"`

	CvUrl string `gorm:"type:varchar(255)" json:"cv_url"`
}

type Business struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `gorm:"not null;unique" json:"user_id"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`

	CompanyName string `gorm:"type:varchar(150);not null" json:"company_name"`

	TaxCode string `gorm:"type:varchar(50);unique" json:"tax_code"`

	Phone string `gorm:"type:varchar(20)" json:"phone"`

	Address string `gorm:"type:varchar(255)" json:"address"`

	LogoUrl string `gorm:"type:varchar(255)" json:"logo_url"`

	IsVerified bool `gorm:"default:false" json:"is_verified"`
}

// 4. Bảng Tin tuyển dụng (Khớp chuẩn chỉnh với DBML của bạn)
type Job struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BusinessID  uint      `gorm:"not null" json:"business_id"` // Khóa ngoại nối sang Business
Business Business `gorm:"foreignKey:BusinessID" json:"business"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Location    string    `gorm:"type:varchar(255);not null" json:"location"`
	Salary      float64   `gorm:"type:decimal(15,2);not null" json:"salary"`
	Slots       int       `gorm:"not null;default:1" json:"slots"`
	Status      string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, closed
	WorkingDate string    `gorm:"type:varchar(100)" json:"working_date"`
	Category    string    `json:"category"`     // Ngành nghề (e.g., "IT", "Marketing")
    JobType     string    `json:"job_type"`
}

type Application struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    JobID     uint      `gorm:"not null" json:"job_id"`     
    StudentID uint      `gorm:"not null" json:"student_id"` 
    Status    string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // Sửa thành pending cho khớp Nuxt
    CoverNote string    `gorm:"type:text" json:"cover_note"`                      // 🌟 Thêm trường Cover Note
    Job       Job       `gorm:"foreignKey:JobID" json:"job,omitempty"`
    Student   Student   `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	OfferSalary    string `json:"offer_salary"`    // Mức lương offer cho sinh viên
    OfferStartDate string `json:"offer_start_date"`// Ngày bắt đầu làm việc
    OfferMessage   string `json:"offer_message"`
}

// Struct bóc tách dữ liệu body gửi lên
type ApplyJobInput struct {
    JobID     uint   `json:"job_id"`
    CoverNote string `json:"cover_note"` // 🌟 Đón đầu Cover Note từ Nuxt 4 gửi lên
}

