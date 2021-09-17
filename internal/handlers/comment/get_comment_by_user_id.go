package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tung.gallery/internal/pkg"
)

func (h *commentHandler) GetCommentListByUserID(c *gin.Context) {
	userIDString := c.Query("id")
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "comment id must be integer value")
		return
	}

	comments, err := h.CommentService.GetCommentsByUserID(c.Request.Context(), userID)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, comments)
}
