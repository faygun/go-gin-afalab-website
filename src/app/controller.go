package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	r.Any("/", func(c *gin.Context) {
		// r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/includes/_layout.html", "templates/views/index.html")))
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.Static("/public", "./public")

	return r
}
