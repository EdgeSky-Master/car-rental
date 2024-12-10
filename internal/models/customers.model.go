package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	// CustomerID string    `gorm:"primaryKey"`
	Name       string    `gorm:"size:100;not null"`
	Email      string    `gorm:"unique;not null"`
	Phone      string    `gorm:"size:50;not null;unique"`
	Birth_date time.Time `gorm:"size:50;type:date"`
	Address string `gorm:"type:text"`
	Drivers_license string `gorm:"size:50"`
	gorm.Model
}