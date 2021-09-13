package pkg

import (
	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/dto"
)

func ResponseSuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	baseResponse := dto.BaseResponse{
		Success:  true,
		DataMsg:  dto.DataMsg{
			Data: data,
		},
	}
	c.JSON(statusCode, baseResponse)
}

func ResponseErrorJSON (c *gin.Context, statusCode int, errors interface{}) {
	baseResponse := dto.BaseResponse{
		Success:  false,
		ErrorMsg: dto.ErrorMsg{
			Error: errors,
		},
	}
	c.JSON(statusCode, baseResponse)
}

