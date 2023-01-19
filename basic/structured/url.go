package structured

import "github.com/gin-gonic/gin"

func register(r gin.IRouter) {
	r.GET("/version", getVersion)
	r.GET("/", getInfo)
	r.POST("/", getInfo)
	r.PUT("/", getInfo)
	r.PATCH("/", getInfo)
	r.DELETE("/", getInfo)

	// 可以把分组拆出去
	//group1 := r.Group("/api/v1")
	//group1.GET("")
	registerApi(r)
}

func registerApi(r gin.IRouter) {
	group1 := r.Group("/api/v1")
	group1.GET("/version", getApiVersion)
}

func (s *Server) register() {
	//s.GET("", nil)
	//group1 := s.Group("/api/v1")
	//group1.GET("")
	register(s)

}
