package user_dto

type UserCreateResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
