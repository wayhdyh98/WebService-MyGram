package routers

import (
	"myGram/controllers"
	"myGram/middlewares"
	"net/http"

	_ "myGram/docs"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @Title MyGram API
// @version 4.0
// @description This is a simple service for managing MyGram
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @License.name Apache 2.0
// @License.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host webservice-mygram-production.up.railway.app
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Simple WebService called myGram. Have you seen the bugs? where? tell me if you found it! :)")
	})

	r.Use(middlewares.CorsMiddleware())

	userRouter := r.Group("/users")
	{
		// Create
		userRouter.POST("/register", controllers.UserRegister)
		// Read
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		// Read All
		photoRouter.GET("/", controllers.GetAllPhoto)
		// Read
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		// Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// Update
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// Delete
		photoRouter.DELETE(("/:photoId"), middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Read All
		commentRouter.GET("/", controllers.GetAllComment)
		// Read
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		// Create
		commentRouter.POST("/", controllers.CreateComment)

		// Update
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// Delete
		commentRouter.DELETE(("/:commentId"), middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	mediaRouter := r.Group("/medias")
	{
		mediaRouter.Use(middlewares.Authentication())
		// Read All
		mediaRouter.GET("/", controllers.GetAllMedia)
		// Read
		mediaRouter.GET("/:mediaId", controllers.GetMediaById)
		// Create
		mediaRouter.POST("/", controllers.CreateMedia)

		// Update
		mediaRouter.PUT("/:mediaId", middlewares.MediaAuthorization(), controllers.UpdateMedia)
		// Delete
		mediaRouter.DELETE(("/:mediaId"), middlewares.MediaAuthorization(), controllers.DeleteMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
