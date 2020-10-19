package sgee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type (
	RouteGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouteGroup
		engine      *Engine
	}

	Engine struct {
		router *router
		*RouteGroup
		groups []*RouteGroup
	}
)

func NewEngine() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouteGroup = &RouteGroup{engine: engine}
	return engine
}

func Default() *Engine {
	engine := NewEngine()
	engine.Use(Logger())
	return engine
}

func (r *RouteGroup) Group(prefix string) *RouteGroup {
	engine := r.engine
	newGroup := &RouteGroup{prefix: r.prefix + prefix, parent: r, engine: engine}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouteGroup) Use(h HandlerFunc) {
	r.middlewares = append(r.middlewares, h)
}

func (r *RouteGroup) Get(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouteGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	routerKey := method + "-" + r.prefix + comp
	r.engine.router.handlers[routerKey] = handlerFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var middlewares []HandlerFunc

	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := NewContext(w, r)
	c.handlers = middlewares
	e.router.handle(c)
}
