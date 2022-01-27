package models

type Store struct {
	Name         string  `json:"name"`
	Status       string  `gorm:"default:ACTIVE" json:"status"`
	Country      string  `json:"country"`
	City         string  `json:"city"`
	Street       string  `json:"street"`
	StreetNumber string  `json:"street_number"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	PostalCode   string  `json:"postal_code"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	InternalId   string  `json:"internal_id"`
}
