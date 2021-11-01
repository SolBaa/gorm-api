package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID         int
	Products   []Product
	TotalPrice float64
}

type Product struct {
	CartID    int
	ProductID string
	Name      string
	Price     string
	Amount    string
}
