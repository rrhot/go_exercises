package main

import (
	gee "sgee"
)

//func onlyForV2() gee.HandlerFunc {
//	return func(c *gee.Context) {
//		// Start timer
//		t := time.Now()
//		// if a server error occurred
//		c.Fail(500, "Internal Server Error")
//		// Calculate resolution time
//		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
//	}
//}
//
//func Gee() {
//	r := gee.New()
//	r.Use(gee.Logger()) // global middleware
//	r.GET("/", func(c *gee.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	})
//
//	v2 := r.Group("/v2")
//	v2.Use(onlyForV2()) // v2 group middleware
//	{
//		v2.GET("/hello/:name", func(c *gee.Context) {
//			// expect /hello/geek tutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		})
//	}
//
//	_ = r.Run(":9999")
//}

func SGee() {
	g := gee.NewEngine()

	g.Get("/index", func(c *gee.Context) {
		c.JSON(200, "/index")
	})

	g.Get("/path", func(c *gee.Context) {
		c.JSON(200, "path")
	})

	_ = g.Run(":9000")
}

func main() {
	SGee()
}
