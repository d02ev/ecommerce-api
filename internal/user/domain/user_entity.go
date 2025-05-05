package domain

import "time"

type UserEntity struct {
	ID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	Email string
	Role uint
	PasswordHash string
	RefreshToken *string

	Addresses []AddressEntity
}
