package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate will set a UUID rather than numeric ID.

func (c *Image) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

// Image -
type Image struct {
	ID uuid.UUID `sql:"type:uuid;primary_key;default:uuid.NewV4().String()" gorm:"index"`
	gorm.Model
	ImageType string `form:"image_type" json:"image_type" `
	UserID    string `form:"user_id" json:"user_id" `
	ItemID    string `form:"item_id" json:"item_id" `
	URL       string `form:"url" json:"url" `
	ThumbURL  string `form:"thumb_url" json:"thumb_url" `
	Path      string `form:"path" json:"path" `
	Position  int    `form:"position" json:"position" `
}
