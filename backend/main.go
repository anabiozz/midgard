package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func port() string{
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func main() {
	engine := gin.Default()

	engine.GET("/api/")

	engine.Run(port())
}
