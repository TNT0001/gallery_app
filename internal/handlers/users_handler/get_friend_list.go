package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (u *userHandler) GetUserFriendList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "error when parse request")
		return
	}

	res, err := u.Services.GetFriendList(uint(id))
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
