package domain

import (
	"time"
)

type CartEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID uint

	CartItems []CartItemEntity
}
