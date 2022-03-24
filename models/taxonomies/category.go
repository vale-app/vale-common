package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SubCategory struct {
	gorm.Model
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	CategoryId uint     `json:"category_id"`
	Category   Category `json:"category"`
}

type Segment struct {
	gorm.Model
	Id            uint        `json:"id"`
	Name          string      `json:"name"`
	SubCategoryId uint        `json:"sub_category_id"`
	SubCategory   SubCategory `json:"sub_category"`
}

type SubSegment struct {
	gorm.Model
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	SegmentId uint    `json:"segment_id"`
	Segment   Segment `json:"segment"`
}
