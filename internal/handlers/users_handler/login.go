package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/pkg"
)

func (u *userHandler) Login(c *gin.Context) {
	req := &userdto.UserLoginRequest{}
	err := c.ShouldBind(req)

	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := u.Services.Login(req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
