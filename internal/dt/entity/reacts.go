package entity

import "gorm.io/gorm"

var reactsTableName = "reacts"

type React struct {
	TypeID  int64 `gorm:"column:type_id"`
	UserID  int64 `gorm:"column:user_id"`
	ImageID int64 `gorm:"column:image_id"`
	gorm.Model
}

// TableName func
func (e *React) TableName() string {
	return reactsTableName
}

var reactTypeTableName = "react_types"

type ReactType struct {
	Type string `gorm:"column:type"`
	gorm.Model
}

// TableName func
func (e *ReactType) TableName() string {
	return reactTypeTableName
}
