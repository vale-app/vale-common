package models

type Discount struct {
	Id             int64  `json:"id"`
	StoreProductId int    `json:"store_product_id"`
	StoreId        int    `json:"store_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Disclaimer     string `json:"disclaimer"`
	Percent        int    `json:"percent"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Status         string `json:"status"`
}
