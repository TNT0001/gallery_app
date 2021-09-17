package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/commentdto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *commentHandler) EditComment(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	req := &commentdto.CommentEditRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}

	err = h.CommentService.Edit(c.Request.Context(), int64(user.ID), req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, nil)
}
