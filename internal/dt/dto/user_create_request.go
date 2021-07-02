package dto

type UserCreateRequest struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
