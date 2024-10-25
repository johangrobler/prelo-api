package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;"`
	Name string    `json:"name"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	category.ID = uuid.New()
	return
}
