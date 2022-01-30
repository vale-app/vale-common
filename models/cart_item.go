package models

import "github.com/kamva/mgm/v3"

type CartItem struct {
	mgm.DefaultModel `bson:",inline"`
	CartId           string  `json:"cart_id" bson:"cart_id"`
	StoreProductId   int64   `json:"store_product_id"`
	Quantity         int64   `json:"quantity"`
	Price            float64 `json:"price"`
}
