package dto

import "time"

type UserUpdateRequest struct {
	Username string     `form:"username"`
	Email    string     `form:"email"`
	Password string     `form:"password"`
	Age      int        `form:"age"`
	ImageURL string     `form:"image_url"`
	Birthday *time.Time `form:"birthday"`
}
