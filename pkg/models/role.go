package models

import "time"

type Role struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	UserID uint `gorm:"column:user_id;not null;index" json:"user_id"`
	Name string `gorm:"column:name;not null;uniqueIndex;size:100" json:"name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
}