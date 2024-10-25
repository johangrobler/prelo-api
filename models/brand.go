package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;"`
	Name string    `json:"name"`
}

func (brand *Brand) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	brand.ID = uuid.New()
	return
}
