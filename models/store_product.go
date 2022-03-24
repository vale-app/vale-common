package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type StoreProduct struct {
	gorm.Model
	StoreId        uint    `json:"store_id"`
	Category       string  `json:"category"`
	SubCategory    string  `json:"sub_category"`
	ProductId      string  `json:"product_id"`
	AvailableUnits int     `json:"available_units"`
	ReservedUnits  int     `json:"reserved_units"`
	Price          float32 `json:"price"`
	Status         string  `json:"status" gorm:"default:'ACTIVE'"`
	InternalId     string  `json:"internal_id"`
	IsCustom       bool    `json:"is_custom" gorm:"default:false"`
	Ean            string  `json:"ean"`
	Sku            string  `json:"sku"`
	ProductName    string  `json:"product_name"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *StoreProduct) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
