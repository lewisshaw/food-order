package routes

import (
	"encoding/json"
	"food_order/app/storage"
	"net/http"
	"strconv"
)

type GetOrderItemResponseJson struct {
	Id    int    `json:"id`
	Name  string `json:"name"`
	Price string `json: "price"`
}
type GetOrderResponseJson struct {
	Id    int                        `json:"id"`
	Email string                     `json:"email"`
	Items []GetOrderItemResponseJson `json:"items"`
}
type GetOrdersResponseJson struct {
	Orders []GetOrderResponseJson `json:"orders"`
}

func GetAllOrdersAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	var ordersResponseJson GetOrdersResponseJson
	for _, order := range storage.GetAllOrders() {
		orderResponseJson := GetOrderResponseJson{Id: int(order.ID), Email: order.Email}
		for _, orderItem := range order.OrderItems {
			priceAsString := strconv.FormatInt(int64(orderItem.Price), 10)
			orderItemResponseJson := GetOrderItemResponseJson{Id: int(orderItem.ID), Name: orderItem.Name, Price: priceAsString}
			orderResponseJson.Items = append(orderResponseJson.Items, orderItemResponseJson)
		}
		ordersResponseJson.Orders = append(ordersResponseJson.Orders, orderResponseJson)
	}

	json.NewEncoder(writer).Encode(ordersResponseJson)
}
