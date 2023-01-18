package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
	"time"
)

func basic(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"message": "pong"})
	})
	r.GET("/message", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, now is %v ", time.Now().Format("2006-01-02 15:04:05"))
	})
}

func handleJson(r *gin.Engine) {
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

func urlRequestHandle(r *gin.Engine) {
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
}

func customResponse(r *gin.Engine) {
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
}

func renderData(r *gin.Engine) {
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `handleJson:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
}

func HandlePost(r *gin.Engine) {
	r.POST("/post/form", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.DefaultPostForm("pwd", "123456")
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"name":     name,
			"password": password,
		})
	})

	// curl -X "POST" "http://localhost:8080/post/form/array" \
	//     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8' \
	//     --data-urlencode "id=1,2" \
	//     --data-urlencode "name=tom,jerry"
	r.POST("/post/form/array", func(c *gin.Context) {
		ids := c.PostFormArray("id")
		names := c.PostFormArray("name")
		if len(ids) != len(names) {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		type user struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}
		data := make([]user, 0, len(ids))
		for i := 0; i < len(ids); i++ {
			data = append(data, user{
				Id:   ids[i],
				Name: names[i],
			})
		}
		c.JSON(http.StatusOK, data)
	})

	// curl -X "POST" "http://localhost:8080/post/form/map" \
	//     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8' \
	//     --data-urlencode "ids[1]=Tom" \
	//     --data-urlencode "ids[2]=Jerry"
	r.POST("/post/form/map", func(c *gin.Context) {
		ids := c.PostFormMap("ids")
		c.JSON(http.StatusOK, ids)
	})

	r.POST("/post/json", func(c *gin.Context) {
		person := &struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{}
		err := c.BindJSON(person)
		if err != nil {
			c.String(http.StatusInternalServerError, "internal server error")
			return
		}
		c.JSON(http.StatusOK, person)
	})

	r.POST("/post/json2", func(c *gin.Context) {
		bs, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		person := &Person{}
		err = json.Unmarshal(bs, person)
		if err != nil {
			c.String(http.StatusInternalServerError, "internal server error")
			return
		}
		c.JSON(http.StatusOK, person)
	})
}

func handleRedirect(r *gin.Engine) {
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ping")
	})

	r.GET("/goto/message", func(c *gin.Context) {
		c.Request.URL.Path = "/message"
		r.HandleContext(c)
	})
}

func defaultHandler(r *gin.Engine) {

}
