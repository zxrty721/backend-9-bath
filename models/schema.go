package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:191"`
	Password string
	Fullname string
	Role     string // admin, staff
	Status   string `gorm:"default:'active'"`
}

type Product struct {
	gorm.Model
	ProductCode  string `gorm:"uniqueIndex;size:191"`
	ProductName  string
	Category     string
	Price        float64
	Quantity     int
	ProductImage string
}

type ProductLog struct {
	gorm.Model
	ProductID   uint
	ProductName string
	Action      string // add, delete, update
	Quantity    int
	UserID      uint
}
