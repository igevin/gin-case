package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"message": "pong"})
	})
	r.GET("/message", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, now is %v ", time.Now().Format("2006-01-02 15:04:05"))
	})
	// 获取URL中的变量
	r.GET("/user/:name/:role", func(c *gin.Context) {
		name := c.Param("name")
		role := c.Param("role")
		c.String(http.StatusOK, "hello, %s, as %s", name, role)
	})
	r.GET("/user2/:name/*role", func(c *gin.Context) {
		name := c.Param("name")
		role := c.Param("role")
		c.String(http.StatusOK, "hello, %s, as %s", name, role)
	})
	// 获取URL参数
	r.GET("param", func(c *gin.Context) {
		name := c.Query("name")
		role := c.Query("role")
		c.JSON(http.StatusOK, gin.H{"name": name, "role": role})
	})
	// URL 字典参数
	// e.g. /map?ids[1]=tom&ids[2]=jerry
	r.GET("/map", func(c *gin.Context) {
		ids, ok := c.GetQueryMap("ids")
		if !ok {
			c.String(http.StatusBadRequest, "no ids param found")
			return
		}
		c.JSON(http.StatusOK, gin.H{"ids": ids})
	})
	r.GET("/map/weak", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		c.JSON(http.StatusOK, ids)
	})
	r.GET("custom1", func(c *gin.Context) {
		c.Header("key1", "value1")
		c.String(http.StatusNotFound, "404 not found")
	})
	r.GET("custom2", func(c *gin.Context) {
		c.Status(http.StatusNotFound)
		c.Header("key1", "value1")
		c.Writer.Header().Add("key2", "value2")
		// writes a header in wire format, not working next line
		//_ = c.Writer.Header().Write(bytes.NewBufferString("key3: value3"))
		_, _ = c.Writer.WriteString("404 not found")
	})
	r.GET("custom3", func(c *gin.Context) {
		c.Header("key1", "value1")
		c.Writer.Header().Add("key2", "value2")
		c.Writer.WriteHeader(http.StatusNotFound)
		_, _ = c.Writer.Write([]byte("404 not found"))
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
