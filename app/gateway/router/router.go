package router

import (
	"github.com/gin-gonic/gin"
	"web3-music-platform/app/gateway/http"
	"web3-music-platform/app/gateway/middleware"
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
			userGroup.POST("/register", http.UserRegisterHandler)
			userGroup.GET("/login", http.UserLoginHandler)
		}
		postGroup := v1.Group("/post")
		postGroup.Use(middleware.JWT())
		{
			postGroup.POST("", http.CreatePostHandler)
			postGroup.GET("/:author_addr", http.GetPostsByAuthorHandler)
		}
		//// 需要登录保护
		//songGroup := v1.Group("/song")
		////authed.Use(middleware.JWT())
		//{
		//	songGroup.GET("/:address", http.GetAllSongsByAuthor)
		//	songGroup.POST("", http.CreateSong)
		//
		//}
	}
	return ginRouter
}
