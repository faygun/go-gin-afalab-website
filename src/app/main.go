package main

func main() {
	r := registerRoutes()

	r.Run(":3000")
}
