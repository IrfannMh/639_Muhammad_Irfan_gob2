package controllers

import (
	"MyGram/config"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.photo_id,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}
func UpdateComment(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}
	CommentUpdate := models.Comment{}
	CommentID := c.Param("commentId")

	if contentType == appJSON {
		c.ShouldBindJSON(&CommentUpdate)
	} else {
		c.ShouldBind(&CommentUpdate)
	}

	err := db.Model(&CommentUpdate).Where("id = ?", CommentID).Updates(&CommentUpdate).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	db.Where("id = ?", CommentID).First(&Comment)

	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.photo_id,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}
func GetComment(c *gin.Context) {

}
func DeleteComment(c *gin.Context) {
	db := config.GetDB()
	Comment := models.Comment{}
	CommentID := c.Param("commentId")

	err := db.Where("id = ?", CommentID).Delete(&Comment).Error
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
