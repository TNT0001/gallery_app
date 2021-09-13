package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/pkg"
)

func (u *userHandler) SignUp(c *gin.Context) {
	req := &user_dto.UserCreateRequest{}
	err := c.ShouldBind(req)

	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err)
		return
	}

	res, err := u.Services.CreateUser(req)

	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
