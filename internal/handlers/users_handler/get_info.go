package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
	"tung.gallery/internal/pkg/mapper"
)

func (u *userHandler) GetUserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "error when parse request")
		return
	}

	user, err := u.Services.FindUserByID(int64(id))
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := mapper.FromUserToUserInfo(user)

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
