package sgee

import "net/http"

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
	}
)

func NewEngine() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

type HandlerFunc func(*Context)

func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
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
