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
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
