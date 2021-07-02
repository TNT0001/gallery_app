package users

import (
	"net/http"

	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/services"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Services services.UserServiceInterface
}

func NewUserHandler(s services.UserServiceInterface) *userHandler {
	return &userHandler{Services: s}
}

func (u *userHandler) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup", gin.H{
		"PageName": "contact",
	})
}

func (u *userHandler) Create(c *gin.Context) {
	req := dto.UserCreateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signup", dto.UserCreateResponse{})
		return
	}

	res, err := u.Services.CreateUser(req)

	if err != nil {
		c.HTML(http.StatusBadRequest, "signup", res)
		return
	}

	c.HTML(http.StatusOK, "signup", res)
}

func (u *userHandler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"PageName": "contact",
	})
}

func (u *userHandler) Login(c *gin.Context) {
	req := dto.UserLoginRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		c.HTML(http.StatusBadRequest, "login", dto.UserLoginResponse{})
		return
	}

	token, err := u.Services.Login(req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login", dto.UserLoginResponse{})
		return
	}
	c.SetCookie("token", token, 60*60*24, "/", "127.0.0.1", true, true)

	c.HTML(http.StatusOK, "home", dto.UserLoginResponse{Login: true})
}

func (u *userHandler) Update(c *gin.Context) {
	req := dto.UserUpdateRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.UserUpdateResponse{Login: true})
		return
	}

	user, _ := c.Get("user")
	oldUser := user.(entity.Users)

	err = u.Services.UpdateUser(oldUser, req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.UserUpdateResponse{Login: true})
		return
	}

	c.HTML(http.StatusBadRequest, "login", dto.UserUpdateResponse{Login: true})
}

func (u *userHandler) Delete(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.HTML(http.StatusBadRequest, "login", dto.UserUpdateResponse{})
	}
	oldUser, ok := user.(entity.Users)
	if !ok {
		c.HTML(http.StatusBadRequest, "login", dto.UserUpdateResponse{})
		return
	}

	err := u.Services.DeleteUser(oldUser)
	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.UserUpdateResponse{Login: true})
		return
	}

	c.HTML(http.StatusOK, "home", dto.UserUpdateResponse{})
}

func (u *userHandler) LogOut(c *gin.Context) {
	c.Set("user", nil)
	c.HTML(http.StatusOK, "home", struct {
		Login bool
	}{Login: false})
}
