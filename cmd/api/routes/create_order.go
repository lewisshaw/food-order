package routes

import (
	"encoding/json"
	"fmt"
	"food_order/app/process"
	"food_order/app/process/request_dto"
	"net/http"
)

type OrderItemRequestJson struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type OrderRequestJson struct {
	Email string                 `json:"email"`
	Items []OrderItemRequestJson `json:"items"`
}

type CreateOrderResponseJson struct {
	Id int `json:"id"`
}

func CreateOrderAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	var orderRequestJson OrderRequestJson
	err := json.NewDecoder(request.Body).Decode(&orderRequestJson)
	if err != nil {
		fmt.Fprint(writer, fmt.Sprintf(`["Error: %s"]`, err.Error()))
		return
	}

	orderProcessRequest := request_dto.CreateOrderRequest{Email: orderRequestJson.Email}
	for _, orderItemJson := range orderRequestJson.Items {
		orderItemProcessRequest := request_dto.CreateOrderItemRequest{Name: orderItemJson.Name, Price: orderItemJson.Price}
		orderProcessRequest.Items = append(orderProcessRequest.Items, orderItemProcessRequest)
	}
	result := process.CreateOrder(orderProcessRequest)

	json.NewEncoder(writer).Encode(CreateOrderResponseJson{Id: result.Id})
}
