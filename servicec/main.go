package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/getsecretcode", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secretcode": "qwerty123456",
		})
	})

	router.Run(":6702")
}