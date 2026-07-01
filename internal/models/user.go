package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password  string    `gorm:"not null;type:varchar(255)" json:"-"`
	Role      string    `gorm:"type:varchar(20);not null" json:"role"` // admin, business, student
	Status    string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Balance   float64   `gorm:"type:decimal(15,2);default:0.00" json:"balance"` // Khớp với decimal của bạn
	CreatedAt time.Time `json:"created_at"`

	Student  *Student  `json:"student,omitempty"`
	Business *Business `json:"business,omitempty"`
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
