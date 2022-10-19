package controllers

import (
	"MyGram/config"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	photoId := c.Param("photoId")
	PhotoUpdate := models.Photo{}
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	db.Where("id = ?", photoId).First(&PhotoUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":         PhotoUpdate.ID,
		"title":      PhotoUpdate.Title,
		"caption":    PhotoUpdate.Caption,
		"photo_url":  PhotoUpdate.PhotoURL,
		"user_id":    PhotoUpdate.UserID,
		"updated_at": PhotoUpdate.UpdatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	db := config.GetDB()
	Photos := []models.Photo{}

	err := db.Preload("Users").Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}

func DeletePhoto(c *gin.Context) {
	db := config.GetDB()
	Photo := models.Photo{}
	photoID := c.Param("photoId")

	err := db.Where("id = ?", photoID).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been succesfully deleted",
	})
}
