package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Name        string    `gorm:"column:name;not null;size:255" json:"name"`
	Description string    `gorm:"column:description;not null;type:text" json:"description"`
	Price       float64   `gorm:"column:price;not null" json:"price"`
	SKU         string    `gorm:"column:sku;not null;uniqueIndex;size:50" json:"sku"`
	StockQty    uint      `gorm:"column:stock_qty;not null" json:"stock_qty"`

	CategoryID uint     `gorm:"column:category_id;not null;uniqueIndex" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`

	OrderItems []OrderItem `gorm:"foreignKey:ProductID;references:ID" json:"order_items"`
	CartItems  []CartItem  `gorm:"foreignKey:ProductID;references:ID" json:"cart_items"`
}
