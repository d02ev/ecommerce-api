package domain

import (
	"time"

	"github.com/d02ev/ecommerce-api/internal/product/domain"
)

type OrderItemEntity struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Quantity   uint
	TotalPrice float64

	OrderID   uint
	Order     OrderEntity
	ProductID uint
	Product   domain.ProductEntity
}
