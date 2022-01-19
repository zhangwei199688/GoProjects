package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	r *router
}

func New() *Engine {
	return &Engine{r : newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	engine.r.addRouter(method,pattern,handler)
}

//注册一个GET方法
func (engine *Engine) GET (pattern string, handler HandlerFunc) {
	engine.r.addRouter("GET", pattern, handler)
}

//注册一个POST方法
func (engine *Engine) POST (pattern string, handler HandlerFunc) {
	engine.r.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//dispatch中心
func (engine *Engine) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.r.handler(c)
}