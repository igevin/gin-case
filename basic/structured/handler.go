package structured

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getVersion(c *gin.Context) {
	version := "0.1"
	c.String(http.StatusOK, "本系统版本为%s", version)
}

func getInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"fullPath": c.FullPath(),
		"method":   c.Request.Method,
		"handler":  c.HandlerName(),
	})
}
