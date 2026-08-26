package models

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
