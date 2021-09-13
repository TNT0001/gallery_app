package entity

import "gorm.io/gorm"

var reactsTableName = "reacts"

type Reacts struct {
	Like uint64 `gorm:"column:like"`
	Dislike uint64 `gorm:"column:dislike"`
	UserID uint `gorm:"column:user_id"`
	gorm.Model
}

// TableName func
func (e *Reacts) TableName() string {
	return reactsTableName
}

