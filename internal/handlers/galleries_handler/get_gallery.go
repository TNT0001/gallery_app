package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (g *galleryHandler) GetGalleryByID(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	galleryIDString := c.Query("gallery_id")
	galleryID, err := strconv.ParseUint(galleryIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	res, err := g.Services.GetGalleryByID(user.ID, uint(galleryID))
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}

func (g *galleryHandler) GetALlGalleryByUserID(c *gin.Context) {
	currentUser, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	userIDString := c.Query("user_id")
	userID, err := strconv.ParseUint(userIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	res, err := g.Services.GetAllGalleriesByUserID(currentUser.ID, uint(userID))
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
