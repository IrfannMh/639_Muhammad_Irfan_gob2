package routers

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetAllCar)
	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
