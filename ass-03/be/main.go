package main

import (
	"be/config"
	"be/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.StartDB()

	router := gin.Default()

	router.GET("/weather", controllers.GetDataWheather)
	router.PUT("/weather", controllers.UpdateWeather)
	router.Run(":8080")
}
