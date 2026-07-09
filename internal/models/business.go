package models

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
