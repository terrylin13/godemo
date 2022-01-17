package main

import (
	"framework"
	"net/http"
)

func main() {
	r := framework.New()
	r.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *framework.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *framework.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
