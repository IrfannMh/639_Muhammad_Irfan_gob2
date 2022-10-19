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
		userRouter.DELETE("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", middlewares.UserAuthorization(), controllers.GetPhoto)
		photoRouter.POST("/", middlewares.UserAuthorization(), controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.UserAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.UserAuthorization(), controllers.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", middlewares.UserAuthorization(), controllers.GetComment)
		commentRouter.POST("/", middlewares.UserAuthorization(), controllers.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.UserAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.UserAuthorization(), controllers.DeleteComment)
	}
	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.GET("/", middlewares.UserAuthorization(), controllers.GetSosmed)
		sosmedRouter.POST("/", middlewares.UserAuthorization(), controllers.CreateSosmed)
		sosmedRouter.PUT("/:socialMediaId", middlewares.UserAuthorization(), controllers.UpdateSosmed)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.UserAuthorization(), controllers.DeleteSosmed)
	}

	return r
}
