package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)
func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	r.GET("/api/myservice/keepalive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			},
		)
	})
	r.Run()
}
