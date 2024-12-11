package models

import "gorm.io/gorm"

type Cars struct {
	// CarID      string  `gorm:"primaryKey"`
	Brand      string  `gorm:"size:100;not null"`
	PlatNumber string  `gorm:"unique;not null"`
	DailyRate  float64 `gorm:"type:decimal(10,2);not null"`
	Status     string  `gorm:"size:50;not null"`
	Color      string  `gorm:"size:50"`
	Year       int     `gorm:"not null"`
	gorm.Model
	
}
