package dto

import "github.com/d02ev/ecommerce-api/pkg/shared/dto"

type OrderDto struct {
	Status          string
	ShippingAddress dto.AddressDto
}
