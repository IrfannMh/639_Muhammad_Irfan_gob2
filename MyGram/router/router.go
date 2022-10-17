package router

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.UpdateUser)
		// userRouter.DELETE("/", middlewares.Authentication(), middlewares.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/")
		// photoRouter.POST("/", middlewares.aut)
		photoRouter.PUT("/")
		photoRouter.DELETE("/")
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/")
		commentRouter.POST("/")
		commentRouter.PUT("/")
		commentRouter.DELETE("/")
	}
	sosmedRouter := r.Group("/socialmedia")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.GET("/")
		sosmedRouter.POST("/")
		sosmedRouter.PUT("/")
		sosmedRouter.DELETE("/")
	}

	return r
}
