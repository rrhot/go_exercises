package main

import "net/http"

func main() {
	r := Default()
	r.GET("/", func(c *Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	v1 := r.Group("/init")
	v1.GET("/index", func(context *Context) {
		context.JSON(10010, "/init/index")
	})

	_ = r.Run(":9999")
}
