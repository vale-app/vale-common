package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartId         int      `json:"cart_id" bson:"cart_id"`
	StoreProductId int      `json:"store_product_id"`
	ProductId      string   `json:"product_id"`
	Quantity       int      `json:"quantity"`
	Price          float64  `json:"price"`
	Product        *Product `json:"product"`
}
