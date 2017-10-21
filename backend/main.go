package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"heimdallProject/midgard/backend/endpoints"
	"heimdallProject/midgard/backend/handlers"
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

	engine.GET("/query/elastic", handlers.ElasticQuery())
	engine.GET("/api/get-charts", endpoints.GetChartsEndpoint())

	engine.Run(port())
}
