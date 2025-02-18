package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	// Create a Gin router
	r := gin.Default()

	dir, _ := os.Getwd()
	// Serve static files from the "static" directory
	r.StaticFS("/", http.Dir(dir))

	// Run the server on port 8080
	r.Run()
}
