package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu

(4)
$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/he/hh", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "/he/hh %s, you're at %s\n", c.Query("test"), c.Path)
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/assets/cssx", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"cssx": c.Param("cssx")})
	})
	r.GET("/assets/ccc/cssx", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"ccc/cssx": c.Param("cssx")})
	})
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.GET("/hello/:xx", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "is /hello/:xx hello %s, you're at %s\n", c.Param("xx"), c.Path)
	})
	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello 1 %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/hello/:name/info", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello 2 %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.GET("/hello/:name/:info", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello 3 %s info: %s,  you're at %s\n", c.Param("name"), c.Param("info"), c.Path)
	})

	r.Run(":9999")
}
