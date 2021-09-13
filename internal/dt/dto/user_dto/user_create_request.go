package user_dto

type UserCreateRequest struct {
	Username string `json:"username" binding:"max=255,required"`
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"min=8,max=24,required"`
}
