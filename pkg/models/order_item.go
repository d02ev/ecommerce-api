package models

import "time"

type OrderItem struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Quantity   uint      `gorm:"column:quantity;not null" json:"quantity"`
	TotalPrice float64   `gorm:"column:total_price;not null" json:"total_price"`

	OrderID   uint    `gorm:"column:order_id;not null;index" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	ProductID uint    `gorm:"column:product_id;not null;index" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
}
