package entity

import (
	"gorm.io/gorm"
)

// GalleriesTableName TableName
var GalleriesTableName = "galleries"

// Galleries struct
type Galleries struct {
	UserID uint   `gorm:"column:user_id; not_null; index"`
	Title  string `gorm:"column:title; not_null"`
	gorm.Model
}

// TableName func
func (g *Galleries) TableName() string {
	return GalleriesTableName
}
