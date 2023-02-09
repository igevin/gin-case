package graceful

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	engine := createEngine()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	StartWithGracefulShutdown(srv)

}

func createEngine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.String(http.StatusOK, "Gin Server started")
	})
	return engine
}
