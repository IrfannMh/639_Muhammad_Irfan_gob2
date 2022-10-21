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

type CommentGet struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserComment
	Photo     PhotoComment
}
type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id`
}

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
		"photo_id":   Comment.PhotoID,
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
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}
func GetComment(c *gin.Context) {
	db := config.GetDB()
	Comments := []models.Comment{}

	err := db.Preload("User").Preload("Photo").Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": err.Error(),
		})
		return
	}
	if len(Comments) < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Failed",
			"Message": "Data not found",
		})
		return
	}
	responses := []CommentGet{}
	for i := 0; i < len(Comments); i++ {
		temp := CommentGet{
			ID:        Comments[i].ID,
			Message:   Comments[i].Message,
			PhotoID:   Comments[i].PhotoID,
			UserID:    Comments[i].UserID,
			CreatedAt: Comments[i].CreatedAt,
			UpdatedAt: Comments[i].UpdatedAt,
			User: UserComment{
				ID:       Comments[i].User.ID,
				Email:    Comments[i].User.Email,
				Username: Comments[i].User.Username,
			},
			Photo: PhotoComment{
				ID:       Comments[i].Photo.ID,
				Title:    Comments[i].Photo.Title,
				Caption:  Comments[i].Photo.Caption,
				PhotoURL: Comments[i].Photo.PhotoURL,
				UserID:   Comments[i].Photo.UserID,
			},
		}
		responses = append(responses, temp)
	}
	c.JSON(http.StatusOK, responses)
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
