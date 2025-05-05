package models

import "time"

type Address struct {
		ID uint `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
		CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
		UserID uint `gorm:"column:user_id;not null;index" json:"user_id"`
		AddLine1 string `gorm:"column:add_line1;not null;size:255" json:"add_line1"`
		AddLine2 *string `gorm:"column:add_line2;size:255" json:"add_line2"`
		Landmark *string `gorm:"column:landmark;size:255" json:"landmark"`
		City string `gorm:"column:city;not null;size:100" json:"city"`
		State string `gorm:"column:state;not null;size:100" json:"state"`
		Country string `gorm:"column:country;not null;size:100" json:"country"`
		ZipCode string `gorm:"column:zip_code;not null;size:20" json:"zip_code"`

		Orders []Order `gorm:"foreignKey:ShippingAddressID;references:ID" json:"orders"`
}