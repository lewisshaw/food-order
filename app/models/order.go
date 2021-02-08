package models

type Order struct {
	gorm.Model
	OrderItems []OrderItem
}
