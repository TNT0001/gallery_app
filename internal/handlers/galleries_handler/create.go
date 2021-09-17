package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/gallerydto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *galleryHandler) CreateGallery(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	req := &gallerydto.GalleryCreateRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}

	res, err := h.Services.CreateGallery(int64(user.ID), req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
