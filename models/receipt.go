package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	SaleId     string  `json:"sale_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	StoreId    string  `json:"store_id"`
	CompanyId  string  `json:"company_id"`
	InternalId string  `json:"internal_id"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Receipt) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
