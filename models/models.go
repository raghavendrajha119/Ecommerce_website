package models

import "gorm.io/gorm"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
type Register struct {
	Name     string
	Email    string
	Password string
}
type Product struct {
	gorm.Model
	Title       string
	Price       float64
	Description string
	Category    string
	Rating      struct {
		Rate  float64
		Count int
	}
}

type ProductCart struct {
	gorm.Model
	ProductID int
	UserID    int
}
