package domain

import (
	"time"
)

type CategoryEntity struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
}
