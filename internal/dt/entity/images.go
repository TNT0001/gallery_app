package entity

import "gorm.io/gorm"

var imagesTableName = "images"

type Images struct {
	GalleryID uint   `gorm:"column:gallery_id"`
	ImageUrl  string `gorm:"column:image_url; not_null; index"`
	Title    string `gorm:"column:title"`
	gorm.Model
}

// TableName func
func (e *Images) TableName() string {
	return imagesTableName
}
