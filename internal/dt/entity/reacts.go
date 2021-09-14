package entity

import "gorm.io/gorm"

var reactsTableName = "reacts"

type React struct {
	TypeID uint64 `gorm:"column:type_id"`
	UserID uint `gorm:"column:user_id"`
	ImageID uint `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *React) TableName() string {
	return reactsTableName
}

var reactTypeTableName = "reacts"

type ReactType struct {
	Type string `gorm:"column:type"`
	gorm.Model
}

// TableName func
func (e *ReactType) TableName() string {
	return reactTypeTableName
}
