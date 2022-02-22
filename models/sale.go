package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	CartId             string  `gorm:"column:cart_id" json:"cart_id"`
	StoreId            string  `gorm:"column:store_id" json:"store_id"`
	Amount             float32 `gorm:"column:amount" json:"amount"`
	Currency           string  `gorm:"column:currency" json:"currency"`
	Discount           float32 `gorm:"column:discount" json:"discount"`
	DiscountPercentage float32 `gorm:"column:discount_percentage" json:"discount_percentage"`
	PaymentMethod      string  `gorm:"column:payment_method" json:"payment_method"`
	SellerId           string  `gorm:"column:seller_id" json:"seller_id"`
	IssueReceipt       bool    `gorm:"column:issue_receipt" json:"issue_receipt"`
	ReceiptId          string  `gorm:"column:receipt_id" json:"receipt_id"`
	Status             string  `gorm:"column:status" json:"status"`
	InternalId         string  `gorm:"column:internal_id" json:"internal_id"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Sale) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
