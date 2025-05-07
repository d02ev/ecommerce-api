package domain

import (
	"time"
)

type OrderEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string

	UserID            uint
	ShippingAddressID uint
}
