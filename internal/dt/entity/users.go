package entity

import (
	"time"

	"gorm.io/gorm"
)

// UsersTableName TableName
var UsersTableName = "users"

// Users struct
type Users struct {
	ID       uint       `gorm:"column:id;primary_key;type:int(11);not null"`
	Username string     `gorm:"column:username;not null"`
	Email    string     `gorm:"column:email;not null"`
	Password string     `gorm:"column:password;not null"`
	Age      int        `gorm:"column:age"`
	ImageURL string     `gorm:"column:image_url"`
	Birthday *time.Time `gorm:"column:birthday;default null"`
	gorm.Model
}

// TableName func
func (i *Users) TableName() string {
	return UsersTableName
}
