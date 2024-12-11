package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ROLE_ADMIN = "admin"
	ROLE_STAFF = "staff"
)

type Users struct {
	Username  string    `json:"username" gorm:"unique"`
	Password_hash  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	LastLogin time.Time `json:"last_login" gorm:"default:null;"`
	gorm.Model

}
