package process

import (
	"food_order/app/process/request",
	"food_order/app/process/result"
	"food_order/app/models"
)

func CreateOrder(request.CreateOrderRequest) result.CreateOrderResult {
	order := models.Order{}
	for _, OrderItemRequest := range CreateOrderRequest.Items {
		orderItem := models.OrderItem{Name: OrderItemRequest.Name, Price: OrderItemRequest.Price}
		order.OrderItems = append(order.OrderItems, orderItem)
	}
	data.SaveOrder(&order)
	return CreateOrderResult{success: true}
}
