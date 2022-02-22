package models

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	CartId             string  `gorm:"column:cart_id" json:"cart_id"`
	StoreId            string  `gorm:"column:store_id" json:"store_id"`
	Amount             float32 `gorm:"column:amount" json:"amount"`
	Discount           float32 `gorm:"column:discount" json:"discount"`
	DiscountPercentage float32 `gorm:"column:discount_percentage" json:"discount_percentage"`
	PaymentMethod      string  `gorm:"column:payment_method" json:"payment_method"`
	SellerId           string  `gorm:"column:seller_id" json:"seller_id"`
	IssueReceipt       bool    `gorm:"column:issue_receipt" json:"issue_receipt"`
	InternalId         string  `gorm:"column:internal_id" json:"internal_id"`
}
