package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Name    string
	Price   int
	OrderID int
}
