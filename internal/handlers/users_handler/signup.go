package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/pkg"
)

func (u *userHandler) SignUp(c *gin.Context) {
	req := &userdto.UserCreateRequest{}
	err := c.ShouldBind(req)
	log.Println(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := u.Services.CreateUser(req)

	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
