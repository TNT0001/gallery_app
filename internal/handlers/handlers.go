package handlers

import (
	"net/http"
	"tung.gallery/internal/pkg"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	pkg.ResponseSuccessJSON(c, http.StatusOK, map[string]string{
		"msg": "hello",
	})
}

func Faq(c *gin.Context) {
	pkg.ResponseSuccessJSON(c, http.StatusOK, map[string]string{
		"msg": "This is Faq page",
	})
}

func Contact(c *gin.Context) {
	pkg.ResponseSuccessJSON(c, http.StatusOK, map[string]string{
		"mailto": "support@tungnguyen.com",
	})
}
