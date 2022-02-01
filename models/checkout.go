package models

type Checkout struct {
	CartId          string `json:"cart_id"`
	PaymentMethodId string `json:"payment_method_id"`
}
