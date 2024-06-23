package router

import (
	"github.com/gin-gonic/gin"
	"web3-music-platform/internal/app/gateway/handlers"
	"web3-music-platform/pkg/middleware"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors())
	//store := cookie.NewStore([]byte("something-very-secret"))
	//ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/register", handlers.UserRegister)
			userGroup.GET("/login", handlers.UserLogin)
		}
		postGroup := v1.Group("/post")
		//postGroup.Use(middleware.JWT())
		{
			postGroup.POST("", middleware.JWT(), handlers.CreatePost)
			postGroup.GET("", handlers.GetPosts)
		}
		songGroup := v1.Group("/song")
		{
			songGroup.POST("", middleware.JWT(), handlers.UploadSong)
			songGroup.GET("", handlers.FindSongs)
			songGroup.GET("/:tx_id", handlers.DownloadSong)
			//songGroup.POST("/purchase", middleware.JWT(), handlers.PurchaseSong)

		}
	}
	return ginRouter
}
