package gallerieshandler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/gallery_dto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (g *galleryHandler) CreateGallery(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	req := &gallery_dto.GalleryCreateRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}

	res, err := g.Services.CreateGallery(user.ID, req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, res)
}
