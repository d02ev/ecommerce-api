package domain

import "time"

type AddressEntity struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID uint
	AddLine1 string
	AddLine2 *string
	Landmark *string
	City string
	State string
	Country string
	ZipCode string
}