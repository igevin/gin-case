package custom

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前可以做些处理

		c.Next()

		// 请求后也可以做些处理
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func TimeTaken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		time.Sleep(time.Millisecond)
		c.Next()
		time.Sleep(time.Millisecond)
		latency := time.Since(t)
		log.Println("time taken: ", latency)
	}
}
