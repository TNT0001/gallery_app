package dto

type UserLoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"min=8,max=24,required"`
}
