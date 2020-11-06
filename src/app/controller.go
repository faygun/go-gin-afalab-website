package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/features", func(c *gin.Context) {
		c.HTML(http.StatusOK, "features.html", nil)
	})

	r.GET("/clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients.html", nil)
	})

	r.GET("/contact", func(c *gin.Context) {
		fmt.Println(34)
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	r.Static("/public", "./public")

	return r
}
