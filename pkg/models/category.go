package models

import "time"

type Category struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Name string `gorm:"column:name;not null;size:100;uniqueIndex" json:"name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
	Products []Product `gorm:"foreignKey:CategoryID;references:ID" json:"products"`
}