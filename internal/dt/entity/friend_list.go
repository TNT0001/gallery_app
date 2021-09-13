package entity

import (
	"gorm.io/gorm"
)

// UserIDTableName TableName
var friendsTableName = "friends"

// UserID struct
type Friend struct {
	UserID uint   `gorm:"column:user_id; not_null; index"`
	FriendID  uint `gorm:"column:friend_id; not_null"`
	gorm.Model
}

// TableName func
func (e *Friend) TableName() string {
	return friendsTableName
}
