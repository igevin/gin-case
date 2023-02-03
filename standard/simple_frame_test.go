package standard

import (
	"net"
	"net/http"
	"strings"
)

// HandleFunc 是业务处理逻辑的抽象
type HandleFunc func(*Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 注册路由的核心抽象
	AddRoute(method, path string, handler HandleFunc)
}

type HttpServer struct {
	// 这里需要有个router抽象，存储path和handler
	defaultRouter
}

// 可以再实现个 HttpsServer
//type HttpsServer struct {
//	HttpServer
//
//	CertFile string
//	KeyFile  string
//}
//
//func (m *HttpsServer) Start(addr string) error {
//	return http.ListenAndServeTLS(addr, m.CertFile, m.KeyFile, m)
//}

var _ Server = &HttpServer{}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		w: w,
		r: r,
	}
	// Do something before serve

	h.serve(ctx)

	// Do something after serve
}

func (h *HttpServer) serve(ctx *Context) {
	// 这里可以通过一个router抽象，来处理URL的业务逻辑
	//n, ok := h.findRoute(ctx.r.Method, ctx.r.URL.Path)
	//if !ok {
	//	ctx.w.WriteHeader(http.StatusNotFound)
	//	_, _ = ctx.w.Write([]byte("Not Found"))
	//}
	//n.handler(ctx)
	ctx.w.WriteHeader(http.StatusOK)
	if strings.HasPrefix(ctx.r.URL.Path, "/other") {
		_, _ = ctx.w.Write([]byte("hello, others"))
		return
	}
	_, _ = ctx.w.Write([]byte("hello, world"))
}

func (h *HttpServer) AddRoute(method, path string, handler HandleFunc) {
	// 不同 path 下的不同 http 方法的处理逻辑，封装在 handler 对象中
	// 这里需要通过 router 对象实现
	h.addRoute(method, path, handler)
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

// 以下为router抽象相关业务逻辑

type router interface {
	addRoute(method string, path string, handler HandleFunc)
	findRoute(method string, path string) (*node, bool)
}

type defaultRouter struct {
	// trees 是按照 HTTP 方法来组织的
	// 如 GET => *node
	trees map[string]*node
}

func (d *defaultRouter) addRoute(method string, path string, handler HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultRouter) findRoute(method string, path string) (*node, bool) {
	//TODO implement me
	panic("implement me")
}

// node 代表路由树的节点
type node struct {
	// 当前节点路径
	path string
	// route 到达该节点的完整的路由路径
	route string

	// children 子节点
	children map[string]*node
	// 通配符 * 表达的节点，任意匹配
	starChild *node
	// 路径参数子节点：形式 :param_name
	paramChild *node

	// handler 命中路由之后执行的逻辑
	handler HandleFunc
}
