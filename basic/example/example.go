package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"message": "pong"})
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
