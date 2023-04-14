package routers

import (
	"myGram/controllers"
	"myGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		photoRouter.POST("/", controllers.CreatePhoto)

		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE(("/:photoId"), middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		commentRouter.POST("/", controllers.CreateComment)

		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE(("/:commentId"), middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	mediaRouter := r.Group("/media")
	{
		mediaRouter.Use(middlewares.Authentication())
		mediaRouter.GET("/", controllers.GetAllMedia)
		mediaRouter.GET("/:mediaId", controllers.GetMediaById)
		mediaRouter.POST("/", controllers.CreateMedia)

		mediaRouter.PUT("/:mediaId", middlewares.MediaAuthorization(), controllers.UpdateMedia)
		mediaRouter.DELETE(("/:mediaId"), middlewares.MediaAuthorization(), controllers.DeleteMedia)
	}

	return r
}
