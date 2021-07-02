package dto

type UserCreateResponse struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Login    bool
	Alert    Alert
}
