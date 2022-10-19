package middlewares

import (
	"MyGram/config"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

func UserAuthorizationWithParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		// UserEmail := userData["email"]
		// idParam := c.Param("userID")
		userIdParam, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
		id := uint(userIdParam)
		User := models.User{}

		err := db.Where("id = ?", UserID).First(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doest'n exist",
			})
			return
		}
		if User.ID != id {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		UserID := uint(userData["id"].(float64))
		// UserEmail := userData["email"]
		// idParam := c.Param("userID")
		// userIdParam, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
		// id := uint(userIdParam)
		User := models.User{}

		err := db.Where("id = ?", UserID).First(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data doest'n exist",
			})
			return
		}
		if User.ID != UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
		}
		c.Next()
	}
}
