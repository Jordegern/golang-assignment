package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"uniwise/animals/controllers"
	"uniwise/animals/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/", HomePage)
	router.GET("/ping", PingHandler)

	router.GET("/animals", controllers.ListAnimals)
	router.POST("/animals", controllers.CreateAnimal)
	router.GET("/animals/:id", controllers.GetAnimal)
	router.PUT("/animals/:id", controllers.UpdateAnimal)
	router.DELETE("/animals/:id", controllers.DeleteAnimal)

	router.Run()
}

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "Why limit happy to an hour when cocktails can make it happy all night?")
}
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
