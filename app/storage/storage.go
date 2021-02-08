package storage

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func getConnection() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		db.AutoMigrate(&OrderItem{}, &Order{})
	}
	if err != nil {
		panic(err)
	}

	return db
}

func SaveOrder(order *Order) Order {
	connection := getConnection()
	connection.Create(order)
	return *order
}

func GetAllOrders() []Order {
	connection := getConnection()
	orders := []Order{}
	result := connection.Preload("OrderItems").Find(&orders)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return orders
}
