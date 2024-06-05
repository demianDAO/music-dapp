package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/idl/pb"
)

func CreatePostHandler(ctx *gin.Context) {
	var req pb.PostCreateRequest
	req.AuthorAddr = ctx.GetString("user_address")

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rpc.CreatePost(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func GetPostsByAuthorHandler(ctx *gin.Context) {
	log.Print("GetPostsByAuthorHandler")
	var req pb.GetPostsByAuthorRequest
	authorAddress := ctx.Query("author_address")
	req.AuthorAddr = authorAddress

	res, err := rpc.GetPostsByAuthor(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
