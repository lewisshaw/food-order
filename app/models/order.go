package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderItems []OrderItem
}
