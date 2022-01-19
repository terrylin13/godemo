package main

import (
	"framework"
	"net/http"
)

func main() {
	r := framework.Default()
	r.GET("/index", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *framework.Context) {
			c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		})

		v1.GET("/hello", func(c *framework.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/", func(c *framework.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v2.GET("/hello/:name", func(c *framework.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *framework.Context) {
			c.JSON(http.StatusOK, framework.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
