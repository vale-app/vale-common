package models

import "github.com/kamva/mgm/v3"

type CartItem struct {
	mgm.DefaultModel `bson:",inline"`
	CartId           string  `json:"cart_id" bson:"cart_id"`
	StoreProductId   int     `json:"store_product_id"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`
}
