package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(r gin.IRouter) {
	// 顺序：先进先出
	r.POST("/user", validate(), handle)
}

func handle(c *gin.Context) {
	validated, ok := c.Get("validated")
	//u, ok := func() (*User, bool) {
	//	if ok {
	//		u, ok := validated.(User)
	//		return &u, ok
	//	}
	//	u := &User{}
	//	err := c.BindJSON(u)
	//	return u, err == nil
	//}()

	if !ok {
		c.Writer.WriteHeader(http.StatusBadRequest)
		_, _ = c.Writer.Write([]byte("bad request, not valid"))
		return
	}
	u := validated.(*User)
	c.JSON(http.StatusOK, u)
}

func validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("validating", "true")
		u := &User{}
		err := c.BindJSON(u)
		if err != nil {
			return
		}
		if u.Group != "admin" {
			return
		}
		c.Set("validated", u)
	}
}

type User struct {
	Name string `json:"name"`
	//Age  int    `json:"age"`
	//birthday time.Time
	Group string `json:"group"`
}
