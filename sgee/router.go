package sgee

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
