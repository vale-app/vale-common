package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/vale-app/vale-common/enums"
	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	CartId             int            `gorm:"column:cart_id" json:"cart_id"`
	PaymentMethodId    string         `gorm:"column:payment_method_id" json:"payment_method_id"`
	Amount             float32        `gorm:"column:amount" json:"amount"`
	Currency           enums.Currency `gorm:"column:currency" json:"currency"`
	Discount           float32        `gorm:"column:discount" json:"discount"`
	DiscountPercentage float32        `gorm:"column:discount_percentage" json:"discount_percent"`
	StoreId            int            `gorm:"column:store_id" json:"store_id"`
	InternalId         string         `json:"internal_id"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Checkout) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
