package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go4digital/booknow-api/routes"
)

const port = "8080"

func main() {
	route := gin.Default()

	routes.Routes(route)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(route.Run(":" + port))
}
