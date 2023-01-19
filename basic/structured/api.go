package structured

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getApiVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "0.0.1"})
}
