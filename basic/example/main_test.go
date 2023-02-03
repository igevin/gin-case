package example

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRunServer(t *testing.T) {
	engine := gin.Default()
	startServer(engine)
}

func startServer(r *gin.Engine) {
	//basic(r)
	//handleJson(r)
	//secureJson(r)
	//urlRequestHandle(r)
	//customResponse(r)
	//renderData(r)
	//HandlePost(r)
	//handleRedirect(r)
	//groupUrl(r)
	//defaultHandler(r)

	RegisterRoute(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
