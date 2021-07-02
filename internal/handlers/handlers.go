package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"PageName": "home",
	})
}

func Faq(c *gin.Context) {
	c.HTML(http.StatusOK, "faq", gin.H{
		"PageName": "faq",
	})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact", gin.H{
		"PageName": "contact",
	})
}
