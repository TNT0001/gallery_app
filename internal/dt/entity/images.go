package entity

import "gorm.io/gorm"

var imagesTableName = "images"

type Images struct {
	ImageUrl string `gorm:"column:image_url; not_null; index"`
	Title    string `gorm:"column:title"`
	CommentID uint  `gorm:"column:comment_id"`
	ReactId	uint 	`gorm:"column:react_id"`
	gorm.Model
}

// TableName func
func (e *Images) TableName() string {
	return imagesTableName
}
