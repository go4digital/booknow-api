package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go4digital/booknow-api/routes"
	"github.com/go4digital/booknow-api/utils"
)

func main() {
	utils.InitLogger()
	port := utils.Getenv("APPLICATION_PORT")
	if port == "" {
		log.Fatal("Application port is empty")
	}
	route := gin.Default()

	routes.Routes(route)

	log.Printf("Server running on http://localhost:%s/", port)
	log.Fatal(route.Run(":" + port))
}
