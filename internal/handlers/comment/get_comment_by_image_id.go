package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *commentHandler) GetCommentListByImageID(c *gin.Context) {
	imageIDString := c.Query("id")
	imageID, err := strconv.ParseInt(imageIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "comment id must be integer value")
		return
	}

	comments, err := h.CommentService.GetCommentsByImageID(c.Request.Context(), imageID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, comments)
}
