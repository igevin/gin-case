package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	basic(r)
	handleJson(r)
	secureJson(r)
	urlRequestHandle(r)
	customResponse(r)
	renderData(r)
	HandlePost(r)
	handleRedirect(r)
	defaultHandler(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
