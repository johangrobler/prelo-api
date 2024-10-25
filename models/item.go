package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name" `
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	BrandID     string    `json:"brand_id"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	item.ID = uuid.New()
	return
}
