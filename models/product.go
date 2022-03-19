package models

import "time"

type Image struct {
	URL   string `json:"name:url" type:"float"`
	Title string `json:"name:title" type:"string"`
}

type Measures struct {
	Height float64 `json:"height" type:"float"`
	Width  float64 `json:"width" type:"float"`
	Depth  float64 `json:"depth" type:"float"`
}

type Product struct {
	ID             string    `json:"id" type:"string"`
	Name           string    `json:"name"`
	Status         string    `json:"status" default:"ACTIVE"`
	Category       string    `json:"category"`
	SubCategory    string    `json:"sub_category"`
	Description    string    `json:"description"`
	ReferencePrice float64   `json:"reference_price" bson:"reference_price"`
	Discontinued   bool      `json:"discontinued" default:"false"`
	InternalId     string    `json:"internal_id"`
	IsPharma       bool      `json:"is_pharma" bson:"is_pharma" default:"false"`
	IsAlcoholic    bool      `json:"is_alcoholic" bson:"is_alcoholic" default:"false"`
	IsCigarette    bool      `json:"is_cigarette" bson:"is_cigarette" default:"false"`
	NeedsRx        bool      `json:"needs_rx" bson:"needs_rx" default:"false"`
	Maker          string    `json:"maker"`
	TradeMark      string    `json:"trade_mark" bson:"trade_mark"`
	Images         []Image   `json:"images"`
	Measures       Measures  `json:"measures"`
	Metadata       string    `json:"metadata"`
	Packing        string    `json:"packing"`
	Weight         float64   `json:"weight"`
	ContainedUnits uint32    `json:"contained_units" bson:"contained_units"`
	Ean            string    `json:"ean" bson:"ean"`
	Sku            string    `json:"sku" bson:"sku"`
	IsForAdult     bool      `json:"is_for_adult" bson:"is_for_adult" default:"false"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

/*func (p *Product) Creating(ctx context.Context) error {
	// Call the DefaultModel Creating hook
	if err := p.DefaultModel.Creating(ctx); err != nil {
		return err
	}

	// Set the default value for the InternalId field
	if p.InternalId == "" {
		newUUID := uuid.NewV4()
		p.InternalId = newUUID.String()
	}

	return nil
}*/
