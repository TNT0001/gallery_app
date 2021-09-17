package entity

import "gorm.io/gorm"

var imagesTableName = "images"

type Images struct {
	UserID    int64  `gorm:"column:user_id"`
	GalleryID int64  `gorm:"column:gallery_id"`
	ImageURL  string `gorm:"column:image_url; not_null; index"`
	Title     string `gorm:"column:title"`
	ImageUUID string `gorm:"column:image_uuid"`
	gorm.Model
}

// TableName func
func (e *Images) TableName() string {
	return imagesTableName
}
