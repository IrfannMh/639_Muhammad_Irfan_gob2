package middlewares

import (
	"JWT/database"
	"JWT/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "Invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data not found",
				"message": "Data doesn.t exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access",
			})
			return
		}
		c.Next()
	}
}
