package models

import "time"

type CartItem struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	CartID uint `gorm:"column:cart_id;not null;index" json:"cart_id"`
	ProductID uint `gorm:"column:product_id;not null;index" json:"product_id"`
	Quantity uint `gorm:"column:quantity;not null" json:"quantity"`
	TotalPrice float64 `gorm:"column:total_price;not null" json:"total_price"`
}