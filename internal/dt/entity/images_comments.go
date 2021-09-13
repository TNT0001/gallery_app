package entity

import "gorm.io/gorm"

var imageCommentsTableName = "image_comments"

type ImageComments struct {
	CommentID uint	`gorm:"column:comment_id"`
	ImageID uint `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *ImageComments) TableName() string {
	return imageCommentsTableName
}
