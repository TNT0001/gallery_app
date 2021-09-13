package entity

import (
	"gorm.io/gorm"
)

// GalleriesTableName TableName
var galleriesTableName = "galleries"

// Galleries struct
type Galleries struct {
	UserID uint   `gorm:"column:user_id; not_null; index"`
	Title  string `gorm:"column:title; not_null"`
	ImageID uint `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *Galleries) TableName() string {
	return galleriesTableName
}
