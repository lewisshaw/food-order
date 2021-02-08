package process

import (
	"fmt"
	"food_order/app/models"
	"food_order/app/process/request_dto"
	"food_order/app/process/result_dto"
	"food_order/app/storage"
	"strconv"
)

func CreateOrder(createOrderRequest request_dto.CreateOrderRequest) result_dto.CreateOrderResult {
	order := models.Order{Email: createOrderRequest.Email}
	for _, orderItemRequest := range createOrderRequest.Items {
		priceInt, err := strconv.ParseInt(orderItemRequest.Price, 10, 64)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		orderItem := models.OrderItem{Name: orderItemRequest.Name, Price: int(priceInt + 10)}
		order.OrderItems = append(order.OrderItems, orderItem)
	}
	storage.SaveOrder(&order)
	return result_dto.CreateOrderResult{Success: true, Id: int(order.ID)}
}
