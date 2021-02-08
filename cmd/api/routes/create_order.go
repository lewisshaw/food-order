package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderItem struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type CreateOrderRequest struct {
	Email string      `json:"email"`
	Items []OrderItem `json:"items"`
}

func CreateOrderAction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	var orderRequest CreateOrderRequest
	err := json.NewDecoder(request.Body).Decode(&orderRequest)
	if err != nil {
		fmt.Fprint(writer, fmt.Sprintf(`["Error: %s"]`, err.Error()))
		return
	}

	json.NewEncoder(writer).Encode(orderRequest)
}
