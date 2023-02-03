package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/igevin/gin-case/basic/example"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	engine := gin.New()
	err := startServer(engine)
	assert.NoError(t, err)
}

func startServer(engine *gin.Engine) error {
	RegisterMetrics()
	engine.Use(promMiddleware())
	registerRoute(engine)
	metrics(engine)

	return engine.Run()

}

func registerRoute(r *gin.Engine) {
	example.RegisterRoute(r)
}
