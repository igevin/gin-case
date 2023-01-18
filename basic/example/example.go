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
	// JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c
	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 如果要按字面对特殊 HTML 字符进行编码，则可以使用 PureJSON。Go 1.6 及更低版本无法使用此功能
	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	secureJson(r)

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

func secureJson(r *gin.Engine) {
	// 可以使用自己的 SecureJSON 前缀，默认预设值为"while(1),"
	prefix := ")(}{][,.\n"
	// prefix 随便设置，这里简化一下
	prefix = ")(}{][,."
	r.SecureJsonPrefix(prefix)

	r.GET("/secure-json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// 这样没有效果
		//c.SecureJSON(http.StatusOK, gin.H{"names": names})
		// 输出 )(}{][,.["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
}
