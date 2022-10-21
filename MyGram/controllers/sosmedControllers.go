package controllers

import (
	"MyGram/config"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaGet struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name" form:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint      `json:"user_id`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           UserSosmed
}

type UserSosmed struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func CreateSosmed(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	Sosmed.UserID = userID
	err := db.Debug().Create(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               Sosmed.ID,
		"name":             Sosmed.Name,
		"social_media_url": Sosmed.SocialMediaURL,
		"user_id":          Sosmed.UserID,
		"created_at":       Sosmed.CreatedAt,
	})
}
func GetSosmed(c *gin.Context) {
	db := config.GetDB()
	Sosmed := []models.SocialMedia{}
	responses := []SocialMediaGet{}

	err := db.Preload("User").Find(&Sosmed).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}
	if len(Sosmed) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Failed",
			"Message": "Data not found",
		})
		return
	}
	for i := 0; i < len(Sosmed); i++ {
		temp := SocialMediaGet{
			ID:             Sosmed[i].ID,
			Name:           Sosmed[i].Name,
			SocialMediaURL: Sosmed[i].SocialMediaURL,
			UserID:         Sosmed[i].UserID,
			CreatedAt:      Sosmed[i].CreatedAt,
			UpdatedAt:      Sosmed[i].UpdatedAt,
			User: UserSosmed{
				ID:       Sosmed[i].User.ID,
				Username: Sosmed[i].User.Username,
			},
		}
		responses = append(responses, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"social medias": responses,
	})
}
func UpdateSosmed(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	Sosmed := models.SocialMedia{}
	SosmedID := c.Param("socialMediaId")
	SosmedUpdate := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Sosmed)
	} else {
		c.ShouldBind(&Sosmed)
	}

	err := db.Model(&Sosmed).Where("id = ?", SosmedID).Updates(&Sosmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}
	db.Where("id = ?", SosmedID).First(&SosmedUpdate)

	c.JSON(http.StatusOK, gin.H{
		"id":               SosmedUpdate.ID,
		"name":             SosmedUpdate.Name,
		"social_media_url": SosmedUpdate.SocialMediaURL,
		"user_id":          SosmedUpdate.UserID,
		"updated_at":       SosmedUpdate.UpdatedAt,
	})
}
func DeleteSosmed(c *gin.Context) {
	db := config.GetDB()
	Sosmed := models.SocialMedia{}
	SosmedID := c.Param("socialMediaId")

	err := db.Where("id = ?", SosmedID).Delete(&Sosmed).Error
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
