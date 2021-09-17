package image

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *imageHandler) GetImageByID(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	imageIDString, ok := c.Params.Get("id")
	if !ok {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "missing param image id")
		return
	}

	imageID, err := strconv.ParseInt(imageIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "image id must be integer value")
		return
	}

	image, err := h.Services.GetImageByID(c.Request.Context(), int64(user.ID), imageID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, image)
}

func (h *imageHandler) GetImageByUserID(c *gin.Context) {
	currentUser, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	userID, err := h.getIDFromParam(c)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
	}

	images, err := h.Services.GetImageByUserID(c.Request.Context(), int64(currentUser.ID), userID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, images)
}

func (h *imageHandler) GetImageByGalleryID(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	galleryIDString, ok := c.Params.Get("id")
	if !ok {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "missing param gallery id")
		return
	}

	galleryID, err := strconv.ParseInt(galleryIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "user id must be integer value")
		return
	}

	images, err := h.Services.GetImageByGalleryID(c.Request.Context(), int64(user.ID), []int64{galleryID})
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, images)
}

func (h *imageHandler) getIDFromParam(c *gin.Context) (int64, error) {
	userIDString, ok := c.Params.Get("id")
	if !ok {
		return 0, errors.New("missing param user id")
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("can't parse id")
	}
	return userID, nil
}
