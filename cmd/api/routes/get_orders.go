package routes

import (
	"encoding/json"
	"food_order/app/storage"
	"food_order/cmd/api/json_structures"
	"net/http"
	"strconv"
)

func GetAllOrdersAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	var ordersResponseJson json_structures.GetOrdersResponseJson
	for _, order := range storage.GetAllOrders() {
		orderResponseJson := json_structures.GetOrderResponseJson{Id: int(order.ID), Email: order.Email}
		for _, orderItem := range order.OrderItems {
			priceAsString := strconv.FormatInt(int64(orderItem.Price), 10)
			orderItemResponseJson := json_structures.GetOrderItemResponseJson{Id: int(orderItem.ID), Name: orderItem.Name, Price: priceAsString}
			orderResponseJson.Items = append(orderResponseJson.Items, orderItemResponseJson)
		}
		ordersResponseJson.Orders = append(ordersResponseJson.Orders, orderResponseJson)
	}

	json.NewEncoder(writer).Encode(ordersResponseJson)
}
