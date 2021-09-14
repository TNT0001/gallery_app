package entity

import "gorm.io/gorm"

var commentsTableName = "comments"

type Comments struct {
	Comment string	`gorm:"column:comment"`
	UserID int64 `gorm:"column:user_id"`
	ImageID int64 `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *Comments) TableName() string {
	return commentsTableName
}
