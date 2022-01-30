package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartId         string  `json:"cart_id" bson:"cart_id"`
	StoreProductId int     `json:"store_product_id"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
}
