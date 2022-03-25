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
	ID              string    `json:"id" type:"string"`
	Name            string    `json:"name"`
	Status          string    `json:"status" default:"ACTIVE"`
	CategoryName    string    `json:"category_name"`
	SubCategoryName string    `json:"sub_category_name"`
	SegmentName     string    `json:"segment_name"`
	SubSegmentName  string    `json:"sub_segment_name"`
	CategoryID      uint      `json:"category_id"`
	SubCategoryID   uint      `json:"sub_category_id"`
	SegmentID       uint      `json:"segment_id"`
	SubSegmentID    uint      `json:"sub_segment_id"`
	Description     string    `json:"description"`
	ReferencePrice  float64   `json:"reference_price"`
	Discontinued    bool      `json:"discontinued" default:"false"`
	InternalId      string    `json:"internal_id"`
	IsPharma        bool      `json:"is_pharma" default:"false"`
	IsAlcoholic     bool      `json:"is_alcoholic" default:"false"`
	IsCigarette     bool      `json:"is_cigarette" default:"false"`
	NeedsRx         bool      `json:"needs_rx" default:"false"`
	Maker           string    `json:"maker"`
	TradeMark       string    `json:"trade_mark"`
	Images          []Image   `json:"images"`
	Measures        Measures  `json:"measures"`
	Metadata        string    `json:"metadata"`
	Packing         string    `json:"packing"`
	Weight          float64   `json:"weight"`
	ContainedUnits  uint32    `json:"contained_units"`
	Ean             string    `json:"ean"`
	Sku             string    `json:"sku"`
	IsForAdult      bool      `json:"is_for_adult" default:"false"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
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
