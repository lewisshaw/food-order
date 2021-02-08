package models

type OrderItem struct {
	gorm.Model
	Name    string
	Price   int
	OrderID int
}
