package dto

type UserLoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
