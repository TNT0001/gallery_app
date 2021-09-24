package react

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *ReactHandler) GetReact(c *gin.Context) {
	reactIDString, ok := c.Params.Get("id")
	if !ok {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "missing param react id")
		return
	}

	reactID, err := strconv.ParseInt(reactIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "react id must be integer value")
		return
	}

	reacts, err := h.Services.GetReactByID(c.Request.Context(), reactID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, reacts)
}
