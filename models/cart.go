package models

import (
	uuid "github.com/satori/go.uuid"
	"github.com/vale-app/vale-common/enums"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId     string           `json:"user_id"`
	StoreId    int              `json:"store_id"`
	Status     enums.CartStatus `json:"status"`
	InternalId string           `json:"internal_id"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID := uuid.NewV4()

	tx.Statement.SetColumn("InternalId", newUUID.String())

	return
}
