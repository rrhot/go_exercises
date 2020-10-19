package sgee

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (r *router) handle(c *Context) {
	n, _ := r.handlers[c.Method+"-"+c.Path]

	if n == nil {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	} else {
		c.handlers = append(c.handlers, n)
	}
	c.Next()
}
