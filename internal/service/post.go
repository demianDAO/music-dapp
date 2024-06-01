package service

import (
	"context"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/repository"
)

// PostService 定义了处理 Post 相关业务逻辑的服务接口
type PostService struct {
	postRepository *repository.PostRepository
}

// NewPostService 创建一个新的 PostService 实例
func NewPostService(postRepository *repository.PostRepository) *PostService {
	return &PostService{postRepository}
}

// CreatePost 创建一个新的 post 记录
func (ps *PostService) CreatePost(ctx context.Context, post *models.Post) error {
	return ps.postRepository.CreatePost(ctx, post)
}

// GetPostsByAuthorAddr 根据作者地址获取其发布的全部 post
func (ps *PostService) GetPostsByAuthorAddr(ctx context.Context, authorAddr string) ([]*models.Post, error) {
	return ps.postRepository.GetPostsByAuthorAddr(ctx, authorAddr)
}
