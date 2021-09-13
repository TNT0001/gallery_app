package users

import (
	"tung.gallery/internal/services/users"
)

type userHandler struct {
	Services users.UserServiceInterface
}

func NewUserHandler(s users.UserServiceInterface) *userHandler {
	return &userHandler{Services: s}
}

//func (u *userHandler) GetSignUpPage(c *gin.Context) {
//	login := utils.CheckLogin(c)
//	c.HTML(http.StatusOK, "signup", dto.BaseResponse{Login: login})
//}



//func (u *userHandler) GetLoginPage(c *gin.Context) {
//	login := utils.CheckLogin(c)
//	baseResponse := dto.BaseResponse{Login: login}
//	c.HTML(http.StatusOK, "login", baseResponse)
//}



//func (u *userHandler) LogOut(c *gin.Context) {
//	baseResponse := dto.BaseResponse{Login: false}
//	c.HTML(http.StatusOK, "home", baseResponse)
//}
