package user_dto

import "time"

type UserGetInfoResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int        `json:"age,omitempty"`
	ImageURL string     `json:"image_url,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
}
