package repository

import (
	"context"
	"gorm.io/gorm"
	"web3-music-platform/internal/api/models"
)

// PostRepository 定义了处理 Post 数据库操作的接口
type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository 创建一个新的 PostRepository 实例
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

// CreatePost 创建一个新的 post 记录
func (pr *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {
	return pr.db.WithContext(ctx).Create(post).Error
}

// GetPostsByAuthorAddr 根据作者地址获取其发布的全部 post
func (pr *PostRepository) GetPostsByAuthorAddr(ctx context.Context, authorAddr string) ([]*models.Post, error) {
	var posts []*models.Post
	if err := pr.db.WithContext(ctx).Where("author_addr = ?", authorAddr).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
