package react

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (h *ReactHandler) CreateReact(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, "fail to get user info")
		return
	}

	req := &reactdto.CreateReactRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}
	req.UserID = int64(user.ID)

	err = h.Services.Create(c.Request.Context(), req)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, nil)
}
