package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type StoreProduct struct {
	gorm.Model
	StoreId        uint     `json:"store_id"`
	CategodyID     uint     `json:"category_id"`
	SubCategoryID  uint     `json:"sub_category_id"`
	SegmentID      uint     `json:"segment_id"`
	SubSegmentID   uint     `json:"sub_segment_id"`
	ProductId      string   `json:"product_id"`
	AvailableUnits int      `json:"available_units"`
	ReservedUnits  int      `json:"reserved_units"`
	Price          float32  `json:"price"`
	Status         string   `json:"status" gorm:"default:'ACTIVE'"`
	InternalId     string   `json:"internal_id"`
	IsCustom       bool     `json:"is_custom" gorm:"default:false"`
	Ean            string   `json:"ean"`
	Sku            string   `json:"sku"`
	ProductName    string   `json:"product_name"`
	Product        *Product `json:"product" gorm:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *StoreProduct) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
