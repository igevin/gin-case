package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igevin/gin-case/advanced/middleware/basic/custom"
	"net/http"
)

func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery(), custom.MyLogger())

	// 你可以为每个路由添加任意数量的中间件。
	r.GET("/benchmark", func(c *gin.Context) {
		c.String(http.StatusOK, "benchmark")
	}, custom.TimeTaken())

	r.GET("/panic", func(c *gin.Context) {
		panic("error")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "this is root")
	})

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := r.Group("/auth")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "succeed to login")
		})
		authorized.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submitted")
		})
		authorized.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read from background")
		})

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			c.String(http.StatusOK, "analysing")
		})
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
