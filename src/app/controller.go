package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"time"

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
	r.GET("/.well-known/acme-challenge", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
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

	r.GET("/services", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", PageSetting{
			CurrentPage: "services",
		})
	})

	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", PageSetting{
			CurrentPage: "contact",
		})
	})

	r.POST("/contact", func(c *gin.Context) {

		email, _ := c.GetPostForm("email")
		username, _ := c.GetPostForm("username")
		message, _ := c.GetPostForm("message")

		date := time.Now().String()

		info := fmt.Sprintf("date: %s \nusername: %s \nemail: %s \nmessage: %s\n\n", date, username, email, message)

		writeText(info)
		c.JSON(http.StatusOK, "Your message has been successfully sent.")
	})

	r.Static("/public", "./public")

	return r
}

func writeText(s string) {
	base := string(http.Dir("public"))

	path := path.Join(base, "file", "email.txt")
	// file, err := os.Stat(path)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	// if file != nil {
	// 	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	// }
	// if file == nil {
	// 	f, err := os.Create(path)
	// }

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(s)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func sendMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
