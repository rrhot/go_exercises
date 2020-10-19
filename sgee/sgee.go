package sgee

import "net/http"

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

func (r *RouteGroup) Group(prefix string) *RouteGroup {
	engine := r.engine
	newGroup := &RouteGroup{prefix: prefix, parent: r, engine: engine}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouteGroup) Get(pattern string, handler HandlerFunc) {
	g.engine.router.handlers[pattern] = handler
}

func (e *Engine) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	routerKey := method + "-" + comp
	e.router.handlers[routerKey] = handlerFunc
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	routerKey := context.Method + "-" + context.Path

	handler := e.router.handlers[routerKey]
	if handler != nil {
		handler(context)
	} else {
		context.String(http.StatusNotFound, "404 NOT FOUND: %s\n", r.URL.Path)
	}
}
