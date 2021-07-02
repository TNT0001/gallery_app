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
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	token, err := u.Services.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.SetCookie("token", token, 60*60*24, "/", "127.0.0.1", true, true)

	c.JSON(http.StatusOK, gin.H{
		"login": "ok",
	})
}

func (u *userHandler) Update(c *gin.Context) {
	req := dto.UserUpdateRequest{}
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user, _ := c.Get("user")
	oldUser := user.(entity.Users)

	err = u.Services.UpdateUser(oldUser, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (u *userHandler) Delete(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	oldUser, ok := user.(entity.Users)
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err := u.Services.DeleteUser(oldUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
