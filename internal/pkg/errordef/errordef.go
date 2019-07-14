package errordef

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorInfo error info
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GinHTTPResponse gin framework http response
func GinHTTPResponse(c *gin.Context, e ErrorInfo) {
	switch e.Code[:3] {
	case "400":
		c.AbortWithStatusJSON(http.StatusBadRequest, e)
	case "500":
		c.AbortWithStatusJSON(http.StatusInternalServerError, e)
	}
}
