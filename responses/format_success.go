package responses

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse(code SuccessResponseCode, data any, c *gin.Context) {
	traceID := c.Request.Header.Get("X-Trace-Id")
	c.Header("X-Trace-Id", traceID)
	c.JSON(code.Status, gin.H{
		"data":    data,
		"message": code.Message,
		"meta": gin.H{
			"code":      code.Code,
			"traceId":   traceID,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	})
}

type SuccessResponseCode struct {
	Code    string
	Message string
	Status  int
}
