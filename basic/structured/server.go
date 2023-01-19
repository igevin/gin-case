package structured

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{
		Engine: engine,
	}
}

func (s *Server) Start(addr string) error {
	// 可以这样
	register(s)
	// 也可以这样，这样更内聚
	//s.register()
	return s.Run(addr)
}
