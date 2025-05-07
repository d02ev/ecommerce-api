package domain

import (
	"time"

	cartDomain "github.com/d02ev/ecommerce-api/internal/cart/domain"
	orderDomain "github.com/d02ev/ecommerce-api/internal/order/domain"
	"github.com/d02ev/ecommerce-api/pkg/shared/entities"
)

type UserEntity struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Email        string
	Role         uint
	PasswordHash string
	RefreshToken *string

	Addresses []entities.AddressEntity
	Orders    []orderDomain.OrderEntity
	Cart      *cartDomain.CartEntity
}
