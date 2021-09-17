package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *commentHandler) GetComment(c *gin.Context) {
	commentIDString := c.Query("id")
	commentID, err := strconv.ParseInt(commentIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "comment id must be integer value")
		return
	}

	comments, err := h.CommentService.GetCommentByID(c.Request.Context(), commentID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, comments)
}
