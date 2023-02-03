package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write([]byte("this is root"))
	})

	registerRoutes(engine)

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
