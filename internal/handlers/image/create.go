package image

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/imagedto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *imageHandler) CreateImage(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	req := &imagedto.ImageUploadRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}
	req.UserID = int64(user.ID)

	imgURL, err := h.Services.CreateImage(c.Request.Context(), req, "no content")

	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, imgURL)
}
