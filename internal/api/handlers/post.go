package handler

import (
	"net/http"
	_type "web3-music-platform/internal/api/handlers/type"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/service"

	"github.com/gin-gonic/gin"
)

// PostHandler 定义了处理 post 请求的 handler
type PostHandler struct {
	postService *service.PostService
}

// NewPostHandler 创建一个新的 PostHandler 实例
func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService}
}

// GetPostsByAuthorAddrHandler 根据作者地址获取其发布的全部 post 的 handler
func (ph *PostHandler) GetPostsByAuthorAddrHandler(c *gin.Context) {
	authorAddr := c.Query("author_addr")
	if authorAddr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing author address"})
		return
	}

	ctx := c.Request.Context()
	posts, err := ph.postService.GetPostsByAuthorAddr(ctx, authorAddr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (ph PostHandler) CreatePostHandler(c *gin.Context) {
	var reqParams _type.PostCreateRequest
	err := c.BindJSON(reqParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		AuthorAddr: reqParams.AuthorAddr,
		Content:    reqParams.Content,
	}

	ctx := c.Request.Context()
	err = ph.postService.CreatePost(ctx, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)

}
