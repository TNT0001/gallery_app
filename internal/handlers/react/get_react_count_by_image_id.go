package react

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *reactHandler) GetReactCountByImageID(c *gin.Context) {
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

	reacts, err := h.Services.GetReactCountByImageID(c.Request.Context(), imageID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, reacts)
}
