package models

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CartId          int     `gorm:"column:cart_id" json:"cart_id"`
	PaymentMethodId string  `gorm:"column:payment_method_id" json:"payment_method_id"`
	Amount          float32 `gorm:"column:amount" json:"amount"`
	Currency        string  `gorm:"column:currency" json:"currency"`
	Discount        float32 `gorm:"column:discount" json:"discount"`
	DiscountPercent float32 `gorm:"column:discount_percent" json:"discount_percent"`
	StoreId         int     `gorm:"column:store_id" json:"store_id"`
}
