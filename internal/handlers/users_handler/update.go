package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (u *userHandler) Update(c *gin.Context) {
	req := &userdto.UserUpdateRequest{}
	err := c.ShouldBind(req)

	log.Println(req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "invalid update form")
		return
	}
	oldUser, err := utils.GetUserFromContext(c)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := u.Services.UpdateUser(oldUser, req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
