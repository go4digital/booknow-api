package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/go4digital/booknow-api/controllers"
	"github.com/go4digital/booknow-api/database"
)

func Routes(router *gin.Engine) {
	db := database.Connect()
	controllers.InitializeDB(db)
	router.GET("/", welcome)
	router.GET("/leads", controllers.GetAllLeads)
	router.POST("/leads", controllers.CreateLead)
	router.GET("/leads/:leadId", controllers.GetLead)
	router.PUT("/leads/:leadId", controllers.UpdateLead)
	router.DELETE("/leads/:leadId", controllers.DeleteLead)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to Book Now API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
}
