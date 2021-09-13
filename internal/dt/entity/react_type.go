package entity

import "gorm.io/gorm"

var reactTypeTableName = "reacts"

type ReactType struct {
	Type string `gorm:"column:type"`
	gorm.Model
}

// TableName func
func (e *ReactType) TableName() string {
	return reactTypeTableName
}


