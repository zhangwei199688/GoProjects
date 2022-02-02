package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}

type Engine struct {
	*RouterGroup
	r *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{r : newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup{
	engine := group.engine
	newGroup := &RouterGroup{
		engine: engine,
		prefix: group.prefix + prefix,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.r.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
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
	var middlewares []HandlerFunc

	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.r.handler(c)
}