package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/dto"
	"tung.gallery/pkg/utils"
)

func Hello(c *gin.Context) {
	login := utils.CheckLogin(c)
	baseResponse := dto.BaseResponse{Login: login}
	c.HTML(http.StatusOK, "home", baseResponse)
}

func Faq(c *gin.Context) {
	login := utils.CheckLogin(c)
	baseResponse := dto.BaseResponse{Login: login}
	c.HTML(http.StatusOK, "faq", baseResponse)
}

func Contact(c *gin.Context) {
	login := utils.CheckLogin(c)
	baseResponse := dto.BaseResponse{Login: login}
	c.HTML(http.StatusOK, "contact", baseResponse)
}
