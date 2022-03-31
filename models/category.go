package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index"`
}

type SubCategory struct {
	gorm.Model
	Id         uint     `json:"id" gorm:"primaryKey"`
	Name       string   `json:"name"`
	CategoryId uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"references:Id"`
}

type Segment struct {
	gorm.Model
	Id            uint        `json:"id" gorm:"primaryKey"`
	Name          string      `json:"name"`
	SubCategoryId uint        `json:"sub_category_id"`
	SubCategory   SubCategory `json:"sub_category" gorm:"references:Id"`
}

type SubSegment struct {
	gorm.Model
	Id        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	SegmentId uint    `json:"segment_id"`
	Segment   Segment `json:"segment" gorm:"references:Id"`
}
