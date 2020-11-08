package main

import (
	"os"
)

func main() {
	r := registerRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
}
