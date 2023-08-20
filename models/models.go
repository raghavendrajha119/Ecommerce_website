package models

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}
type Register struct {
	Name     string
	Email    string
	Password string
}
type Product struct {
	gorm.Model
	ID          uint
	Title       string
	Price       float64
	Description string
	Category    string
	Image       string
	Quantity    int
}
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Quantity  int
}
type BoughtProduct struct {
	gorm.Model
	UserID      uint
	ProductID   uint
	Quantity    int
	TotalAmount float64
}
