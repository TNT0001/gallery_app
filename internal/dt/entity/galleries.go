package entity

import (
	"gorm.io/gorm"
)

// GalleriesTableName TableName
var GalleriesTableName = "galleries"

// Galleries struct
type Galleries struct {
	UserID uint   `gorm:"not_null;index"`
	Title  string `gorm:"not_null"`
	gorm.Model
}

// TableName func
func (g *Galleries) TableName() string {
	return GalleriesTableName
}
