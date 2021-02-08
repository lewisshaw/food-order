package json_structures

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
