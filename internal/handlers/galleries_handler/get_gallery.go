package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *galleryHandler) GetGalleryByID(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	galleryIDString := c.Query("gallery_id")
	galleryID, err := strconv.ParseInt(galleryIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	res, err := h.Services.GetGalleryByID(int64(user.ID), galleryID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}

func (h *galleryHandler) GetALlGalleryByUserID(c *gin.Context) {
	currentUser, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	userIDString := c.Query("user_id")
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	res, err := h.Services.GetAllGalleriesByUserID(int64(currentUser.ID), userID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
