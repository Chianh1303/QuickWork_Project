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
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;unique" json:"user_id"`
	FullName  string    `gorm:"type:varchar(100);not null" json:"full_name"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Gender    string    `gorm:"type:varchar(10)" json:"gender"`      // Mới bổ sung
	AvatarUrl string    `gorm:"type:varchar(255)" json:"avatar_url"`  // Mới bổ sung
	Skills    string    `gorm:"type:text" json:"skills"`
	CvUrl     string    `gorm:"type:varchar(255)" json:"cv_url"`
}

type Business struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;unique" json:"user_id"`
	CompanyName string    `gorm:"type:varchar(150);not null" json:"company_name"`
	TaxCode     string    `gorm:"type:varchar(50);unique" json:"tax_code"`
	Phone       string    `gorm:"type:varchar(20)" json:"phone"`       // Mới bổ sung
	Address     string    `gorm:"type:varchar(255)" json:"address"`   // Mới bổ sung
	LogoUrl     string    `gorm:"type:varchar(255)" json:"logo_url"`   // Mới bổ sung
	IsVerified  bool      `gorm:"default:false" json:"is_verified"`
}

// 4. Bảng Tin tuyển dụng (Khớp chuẩn chỉnh với DBML của bạn)
type Job struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BusinessID  uint      `gorm:"not null" json:"business_id"` // Khóa ngoại nối sang Business
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Location    string    `gorm:"type:varchar(255);not null" json:"location"`
	Salary      float64   `gorm:"type:decimal(15,2);not null" json:"salary"`
	Slots       int       `gorm:"not null;default:1" json:"slots"`
	Status      string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, closed
	WorkingDate string    `gorm:"type:varchar(100)" json:"working_date"`
}

type Application struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	JobID     uint      `gorm:"not null" json:"job_id"`     // Khóa ngoại nối sang bảng Job
	StudentID uint      `gorm:"not null" json:"student_id"` // Khóa ngoại nối sang bảng Student
	Status    string    `gorm:"type:varchar(20);default:'applied'" json:"status"` // applied, approved, rejected
	Job       Job       `gorm:"foreignKey:JobID" json:"job,omitempty"`
	Student   Student   `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}