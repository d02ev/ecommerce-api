package models

import "time"

type Cart struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	UserID uint `gorm:"column:user_id;not null;index" json:"user_id"`

	CartItems []CartItem `gorm:"foreignKey:CartID;references:ID" json:"cart_items"`
}