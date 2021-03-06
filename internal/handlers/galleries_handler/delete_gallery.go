package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *galleryHandler) Delete(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	galleryIDString, ok := c.Params.Get("id")
	if !ok {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "missing gallery id")
		return
	}

	galleryID, err := strconv.ParseInt(galleryIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	res, err := h.Services.Delete(int64(user.ID), galleryID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
