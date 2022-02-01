package models

type CartDetail struct {
	Cart       *Cart       `json:"cart"`
	CartItems  *[]CartItem `json:"cart_items"`
	Amount     float64     `json:"amount"`
	TotalItems int         `json:"total_items"`
}
