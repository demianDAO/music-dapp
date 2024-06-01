package routes

import (
	"github.com/gin-gonic/gin"
	"web3-music-platform/controllers"
	handler "web3-music-platform/internal/api/handlers"
	"web3-music-platform/internal/component/db"
	"web3-music-platform/internal/repository"
	"web3-music-platform/internal/service"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")

	{
		v1.GET("/", controllers.GetUsers)
		v1.GET("/:id", controllers.GetUserByID)
		v1.POST("/", controllers.CreateUser)
		v1.PUT("/:id", controllers.UpdateUser)
		v1.DELETE("/:id", controllers.DeleteUser)
	}

	{
		v1.GET("/", controllers.GetMusic)
		v1.GET("/:id", controllers.GetMusicByID)
		v1.POST("/", controllers.CreateMusic)
		v1.PUT("/:id", controllers.UpdateMusic)
		v1.DELETE("/:id", controllers.DeleteMusic)
	}

	{

		postGroup := v1.Group("/post")
		postHandler := handler.NewPostHandler(service.NewPostService(repository.NewPostRepository(db.DbEngine)))
		postGroup.POST("/", postHandler.CreatePostHandler)
		postGroup.GET("/:author_addr", postHandler.CreatePostHandler)
	}

}
