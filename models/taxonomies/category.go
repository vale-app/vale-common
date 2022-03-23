package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SubCategory struct {
	gorm.Model
	Id       int64    `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

type Segment struct {
	gorm.Model
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	SubCategory SubCategory `json:"sub_category"`
}

type SubSegment struct {
	gorm.Model
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Segment Segment `json:"segment"`
}
