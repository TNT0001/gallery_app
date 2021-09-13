package entity

import "gorm.io/gorm"

var reactsTableName = "reacts"

type Reacts struct {
	TypeID uint64 `gorm:"column:type_id"`
	UserID uint `gorm:"column:user_id"`
	ImageID uint `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *Reacts) TableName() string {
	return reactsTableName
}

