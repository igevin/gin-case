package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igevin/gin-case/basic/structured"
)

func main() {
	s := structured.NewServer(gin.Default())
	err := s.Start(":8080")
	if err != nil {
		panic(err)
	}
}
