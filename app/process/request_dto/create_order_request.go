package request_dto

type CreateOrderItemRequest struct {
	Name  string
	Price string
}

type CreateOrderRequest struct {
	Email string
	Items []CreateOrderItemRequest
}
