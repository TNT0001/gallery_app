package entity

import "gorm.io/gorm"

var imageReactsTableName = "image_reacts"

type ImageReacts struct {
	ReactID uint	`gorm:"column:react_id"`
	ImageID uint `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *ImageReacts) TableName() string {
	return imageReactsTableName
}
