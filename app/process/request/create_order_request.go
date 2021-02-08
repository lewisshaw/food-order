package request

type OrderItem struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type CreateOrderRequest struct {
	Email string      `json:"email"`
	Items []OrderItem `json:"items"`
}
