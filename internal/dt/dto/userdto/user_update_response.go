package userdto

import "time"

type UserUpdateResponse struct {
	Username string     `json:"username,omitempty"`
	Email    string     `json:"email,omitempty"`
	Password string     `json:"password,omitempty"`
	Age      int        `json:"age,omitempty"`
	ImageURL string     `json:"image_url,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}
