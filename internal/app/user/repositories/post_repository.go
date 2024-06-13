package repositories

import (
	"context"
	"gorm.io/gorm"
	"log"
	"web3-music-platform/internal/app/user/db"
	"web3-music-platform/internal/app/user/models"
)

type PostDao struct {
	*gorm.DB
}

func NewPostDao(ctx context.Context) *PostDao {
	return &PostDao{db.NewDBClient(ctx)}
}

func (pd *PostDao) CreatePost(post *models.Post) error {
	return pd.Model(&models.Post{}).Create(&post).Error
}

func (pd *PostDao) GetPosts(userAddr string) ([]*models.Post, error) {
	var posts []*models.Post
	log.Print("userAddr", userAddr)
	err := pd.Model(models.Post{}).Where("user_addr = ?", userAddr).Find(&posts).Error
	log.Print("GetPostsByAuthor", posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
