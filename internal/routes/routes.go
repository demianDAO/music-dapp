package routes

import (
	"github.com/gin-gonic/gin"
	handler "web3-music-platform/internal/api/handlers"
	"web3-music-platform/internal/component/db"
	"web3-music-platform/internal/repository"
	"web3-music-platform/internal/service"
)

func SetupRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		postGroup := v1.Group("/post")
		postHandler := handler.NewPostHandler(service.NewPostService(repository.NewPostRepository(db.DbEngine)))
		postGroup.POST("/", postHandler.CreatePostHandler)
		postGroup.GET("/:author_addr", postHandler.CreatePostHandler)
	}
	{
		songGroup := v1.Group("/song")
		songHandler := handler.NewSongHandler(service.NewSongService(repository.NewSongRepository(db.DbEngine)))
		songGroup.POST("/", songHandler.CreateSongHandler)
		songGroup.GET("/:author_addr", songHandler.GetAllSongByAuthorHandler)
	}
	{
		userGroup := v1.Group("/user")
		userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db.DbEngine)))
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:addr", userHandler.GetUser)
	}

}
