package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageSetting struct {
	PageTitle       string
	PageDescription string
	CurrentPage     string
}

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", PageSetting{
			CurrentPage: "home",
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", PageSetting{
			CurrentPage: "about",
		})
	})

	r.GET("/features", func(c *gin.Context) {
		c.HTML(http.StatusOK, "features.html", PageSetting{
			CurrentPage: "features",
		})
	})

	r.GET("/clients", func(c *gin.Context) {
		c.HTML(http.StatusOK, "clients.html", PageSetting{
			CurrentPage: "clients",
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		fmt.Println(34)
		c.HTML(http.StatusOK, "contact.html", PageSetting{
			CurrentPage: "contact",
		})
	})

	r.Static("/public", "./public")

	return r
}
