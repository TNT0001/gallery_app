package entity

import (
	"time"

	"gorm.io/gorm"
)

// UsersTableName TableName
var usersTableName = "users"

// Users struct
type Users struct {
	Username string     `gorm:"column:username;not null"`
	Email    string     `gorm:"column:email;not null"`
	Password string     `gorm:"column:password;not null"`
	Age      int        `gorm:"column:age"`
	ImageURL string     `gorm:"column:image_url"`
	Birthday *time.Time `gorm:"column:birthday;default null"`
	*gorm.Model
}

// TableName func
func (e *Users) TableName() string {
	return usersTableName
}
