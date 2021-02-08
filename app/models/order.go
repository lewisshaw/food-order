package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Email      string
	OrderItems []OrderItem
}
