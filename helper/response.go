package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPResponse struct {
	Status      bool        `json:"status"`
	ErrorCode   string      `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func ResponseOK(c *gin.Context, data interface{}, description string) {
	c.JSON(http.StatusOK, HTTPResponse{
		Status:      true,
		Description: description,
		Data:        data,
	})
}

func ResponseError(c *gin.Context, errorCode string, description string, httpStatusCode int) {
	c.JSON(httpStatusCode, HTTPResponse{
		Status:      false,
		ErrorCode:   errorCode,
		Description: description,
	})
}
