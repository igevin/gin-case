package standard

import (
	"net"
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 注册路由的核心抽象
	//AddRoute(method, path string, handler HandleFunc)
}

type HttpServer struct {
}

var _ Server = &HttpServer{}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{
		w: w,
		r: r,
	}
	// Do something before serve

	h.serve(ctx)

	// Do something after serve
}

func (h *HttpServer) serve(ctx Context) {
	// 这里可以设计一个router抽象，来处理URL的业务逻辑
	ctx.w.WriteHeader(http.StatusOK)
	if strings.HasPrefix(ctx.r.URL.Path, "/other") {
		_, _ = ctx.w.Write([]byte("hello, others"))
		return
	}
	_, _ = ctx.w.Write([]byte("hello, world"))
}

//func (h *HttpServer) Start(addr string) error {
//	// 这个是阻塞的，要放在最后，没发在它之后做点什么
//	return http.ListenAndServe(addr, h)
//}

func (h *HttpServer) Start(addr string) error {
	// 先启动端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 端口启动后，可以先做些事情，再启动服务
	// 如把服务注册到管理平台或etcd等

	// 这个是阻塞的，要放在最后，没发在它之后做点什么
	return http.Serve(listener, h)
}
