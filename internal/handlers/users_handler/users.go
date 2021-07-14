package users

import (
	"net/http"

	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/services"
	"tung.gallery/pkg/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Services services.UserServiceInterface
}

func NewUserHandler(s services.UserServiceInterface) *userHandler {
	return &userHandler{Services: s}
}

func (u *userHandler) GetSignUpPage(c *gin.Context) {
	login := utils.CheckLogin(c)
	c.HTML(http.StatusOK, "signup", dto.BaseResponse{Login: login})
}

func (u *userHandler) SignUp(c *gin.Context) {
	login := utils.CheckLogin(c)

	req := dto.UserCreateRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		baseResponse := utils.BaseResponse(login, services.AlertLvlInfo, "Invalid create account request")
		c.HTML(http.StatusBadRequest, "signup", dto.UserCreateResponse{
			BaseResponse: baseResponse})
		return
	}

	res, err := u.Services.CreateUser(req)
	res.Login = login

	if err != nil {
		c.HTML(http.StatusBadRequest, "signup", res)
		return
	}

	c.Redirect(http.StatusFound, "/user/login")
}

func (u *userHandler) GetLoginPage(c *gin.Context) {
	login := utils.CheckLogin(c)
	baseResponse := dto.BaseResponse{Login: login}
	c.HTML(http.StatusOK, "login", baseResponse)
}

func (u *userHandler) Login(c *gin.Context) {
	req := dto.UserLoginRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		baseResponse := utils.BaseResponse(false, services.AlertLvlInfo, "invalid email or password")
		c.HTML(http.StatusBadRequest, "login", dto.UserLoginResponse{BaseResponse: baseResponse})
		return
	}

	res, err := u.Services.Login(req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login", res)
		return
	}
	c.SetCookie("token", res.Token, 60*60*24, "/", "localhost", true, true)

	c.Redirect(http.StatusFound, "/gallery")
}

func (u *userHandler) Update(c *gin.Context) {
	req := dto.UserUpdateRequest{}
	err := c.ShouldBind(&req)

	login := utils.CheckLogin(c)

	if err != nil {
		baseResponse := utils.BaseResponse(login, services.AlertLvlInfo, "invalid update form")
		c.HTML(http.StatusBadRequest, "update", dto.UserUpdateResponse{BaseResponse: baseResponse})
		return
	}

	user, _ := c.Get("user")
	oldUser := user.(entity.Users)

	res, err := u.Services.UpdateUser(oldUser, req)
	res.Login = login
	if err != nil {
		c.HTML(http.StatusBadRequest, "update", res)
		return
	}

	c.HTML(http.StatusOK, "update", res)
}

func (u *userHandler) Delete(c *gin.Context) {
	login := utils.CheckLogin(c)
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		baseResponse := utils.BaseResponse(login, services.AlertLvlInfo, utils.ErrUserNotFound.Error())
		c.HTML(http.StatusOK, "update", dto.UserDeleteResponse{BaseResponse: baseResponse})
	}

	res, err := u.Services.DeleteUser(user)
	res.Login = login
	if err != nil {
		c.HTML(http.StatusInternalServerError, "update", res)
		return
	}

	res.Login = false
	c.HTML(http.StatusOK, "home", res)
}

func (u *userHandler) LogOut(c *gin.Context) {
	baseResponse := dto.BaseResponse{Login: false}
	c.HTML(http.StatusOK, "home", baseResponse)
}
