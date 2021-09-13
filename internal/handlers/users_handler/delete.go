package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (u *userHandler) Delete(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError,"invalid update form")
		return
	}

	res, err := u.Services.DeleteUser(user)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError,"invalid update form")
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
