package models

import "time"

type Skill struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null;type:varchar(100)" json:"name"`
	Category  string    `gorm:"type:varchar(100);default:'General'" json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
