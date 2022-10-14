package controllers

import (
	"be/config"
	"be/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// update data
func UpdateWeather(c *gin.Context) {
	db := config.GetDB()
	_ = db
	var (
		weather    models.Weather
		newWeather models.Weather
	)

	json.NewDecoder(c.Request.Body).Decode(&newWeather)

	err := db.Where("id = ?", 1).First(&weather).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  err.Error(),
			"message": "data wheather not found",
		})
		return
	}
	err = db.Model(&weather).Updates(&newWeather).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"result":  err.Error(),
			"message": "cannot update weather",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": weather,
	})
}

// get data
func GetDataWheather(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	db := config.GetDB()
	_ = db
	weather := models.Weather{}

	err := db.Where("id = ?", 1).First(&weather).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result":  err.Error(),
			"message": "data wheather not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"wind":  weather.Wind,
			"water": weather.Water,
		},
	})

}
