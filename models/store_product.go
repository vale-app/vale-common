package models

import "gorm.io/gorm"

type StoreProduct struct {
	gorm.Model
	StoreId        uint    `json:"store_id"`
	Category       string  `json:"category"`
	ProductId      string  `json:"product_id"`
	AvailableUnits int     `json:"available_units"`
	ReservedUnits  int     `json:"reserved_units"`
	Price          float32 `json:"price"`
	Status         string  `json:"status" gorm:"default:'ACTIVE'"`
	InternalId     string  `json:"internal_id"`
	IsCustom       bool    `json:"is_custom" gorm:"default:false"`
}
