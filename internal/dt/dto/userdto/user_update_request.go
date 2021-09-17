package userdto

import "time"

type UserUpdateRequest struct {
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Age      int        `json:"age"`
	ImageURL string     `json:"image_url"`
	Birthday *time.Time `json:"birthday"`
}
