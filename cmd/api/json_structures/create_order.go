package json_structures

type CreateOrderItemRequestJson struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

type CreateOrderRequestJson struct {
	Email string                       `json:"email"`
	Items []CreateOrderItemRequestJson `json:"items"`
}

type CreateOrderResponseJson struct {
	Id int `json:"id"`
}
