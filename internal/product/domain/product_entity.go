package domain

import (
	"time"
)

type ProductEntity struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Price       float64
	SKU         string
	StockQty    uint
	CategoryID  uint
	Category    CategoryEntity
}
