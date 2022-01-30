package models

import (
	"github.com/kamva/mgm/v3"
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	mgm.DefaultModel `bson:",inline"`
	UserId           string `json:"user_id"`
	StoreId          int    `json:"store_id"`
	InternalId       string `json:"internal_id"`
}

func (p *Cart) Creating() error {
	// Call the DefaultModel Creating hook
	if err := p.DefaultModel.Creating(); err != nil {
		return err
	}

	// Set the default value for the InternalId field
	if p.InternalId == "" {
		newUUID := uuid.NewV4()
		p.InternalId = newUUID.String()
	}

	return nil
}
