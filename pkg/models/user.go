package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Name string `gorm:"column:name;not null;size:100" json:"name"`
	Email string `gorm:"column:email;not null;uniqueIndex;size:150" json:"email"`
	Role uint `gorm:"column:role;not null;default:0" json:"role"`
	PasswordHash string `gorm:"column:password_hash;not null;size:255" json:"password_hash"`
	RefreshToken *string `gorm:"column:refresh_token;uniqueIndex;size:255" json:"refresh_token"`

	Cart *Cart `gorm:"foreignKey:UserID;references:ID" json:"cart"`
	Orders []Order `gorm:"foreignKey:UserID;references:ID" json:"orders"`
	Addresses []Address `gorm:"foreignKey:UserID;references:ID" json:"addresses"`
}