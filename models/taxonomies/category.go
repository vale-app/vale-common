package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SubCategory struct {
	gorm.Model
	Id         int64    `json:"id"`
	Name       string   `json:"name"`
	CategoryId int64    `json:"category_id"`
	Category   Category `json:"category"`
}

type Segment struct {
	gorm.Model
	Id            int64       `json:"id"`
	Name          string      `json:"name"`
	SubCategoryId int64       `json:"sub_category_id"`
	SubCategory   SubCategory `json:"sub_category"`
}

type SubSegment struct {
	gorm.Model
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	SegmentId int64   `json:"segment_id"`
	Segment   Segment `json:"segment"`
}
