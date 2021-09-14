package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/dt/dto/gallery_dto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (g *galleryHandler) UpdateGallery(c *gin.Context) {
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

	galleryID, err := strconv.ParseUint(galleryIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "gallery id must be integer value")
		return
	}

	req := &gallery_dto.GalleryUpdateRequest{}
	err = c.ShouldBind(&req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	req.ID = uint(galleryID)

	res, err := g.Services.Update(user.ID, req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
