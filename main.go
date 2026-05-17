package main

import (
	"svc-dox/response"
	"svc-dox/response/codes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/dox/ip", func(c *gin.Context) {
		ip := c.Query("ip")
		if ip == "" {
			response.SendErrorResponse(codes.ErrIpNotProvided, nil, c)
		}

		c.JSON(200, gin.H{})
	})

	r.Run("[::]:80")
}
