package react

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tung.gallery/internal/pkg"
)

func (h *reactHandler) GetReactType(c *gin.Context) {
	reacts, err := h.Services.GetReactMap(c.Request.Context())
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, reacts)
}
