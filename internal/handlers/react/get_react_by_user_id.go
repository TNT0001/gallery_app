package react

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *ReactHandler) GetReactByUserID(c *gin.Context) {
	userIDString, ok := c.Params.Get("id")
	if !ok {
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "missing param user id")
		return
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "user id must be integer value")
		return
	}

	reacts, err := h.Services.GetReactByUserID(c.Request.Context(), userID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, reacts)
}
