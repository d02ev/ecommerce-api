package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Status    string    `gorm:"column:status;not null;size:50;index" json:"status"`

	UserID            uint    `gorm:"column:user_id;not null;index" json:"user_id"`
	User              User    `gorm:"foreignKey:UserID;references:ID" json:"user"`
	ShippingAddressID uint    `gorm:"column:shipping_address_id;not null;index" json:"shipping_address_id"`
	ShippingAddress   Address `gorm:"foreignKey:ShippingAddressID;references:ID" json:"address"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"order_items"`
}
