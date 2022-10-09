package main

import (
	"ass-02/config"
	"ass-02/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrder)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:ordersId", inDB.DeleteOrder)
	router.Run(":3000")
}
