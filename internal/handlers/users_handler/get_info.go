package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/pkg"
)

func (u *userHandler) GetUserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "error when parse request")
		return
	}

	user, err := u.Services.FindUserById(uint(id))
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	res := user_dto.UserGetInfoResponse{
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		ImageURL: user.ImageURL,
		Birthday: user.Birthday,
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
