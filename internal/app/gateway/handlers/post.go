package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"web3-music-platform/internal/app/gateway/services"
	"web3-music-platform/pkg/rdb"
)

func CreatePost(ctx *gin.Context) {
	userAddr := ctx.GetString("user_addr")
	content := ctx.PostForm("content")

	log.Print("CreatePostHandler", userAddr, content)

	go func() {
		post, err := services.CreatePost(ctx, userAddr, content)
		if err != nil {
			log.Print("Failed to create post: ", err.Error())
			return
		}
		log.Print("Adding post to redis", post)
		err = rdb.RedisInstance.AddPost(ctx, post.UserAddr, post)
		if err != nil {
			log.Print("Failed to add post to redis: ", err.Error())
			return
		} else {
			log.Print("Successfully added post to redis")
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
	})
}

func GetPosts(ctx *gin.Context) {
	log.Print("GetPostsHandler")
	userAddr := ctx.Query("user_address")

	posts, err := rdb.RedisInstance.GetPosts(ctx, userAddr)
	log.Print("GetPostsHandler", posts)
	if err == redis.Nil {
		log.Print("Failed to get posts from redis and will query from services")
		posts, err = services.GetPosts(ctx, userAddr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		log.Print("Successfully get posts to redis")

	}
	ctx.JSON(http.StatusOK, posts)
}
