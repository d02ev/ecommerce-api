package domain

import (
	"time"

	"github.com/d02ev/ecommerce-api/internal/product/domain"
)

type CartItemEntity struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Quantity   uint
	TotalPrice float64

	CartID    uint
	Cart      CartEntity
	ProductID uint
	Product   domain.ProductEntity
}
