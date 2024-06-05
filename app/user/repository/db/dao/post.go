package dao

import (
	"context"
	"gorm.io/gorm"
	"log"
	"web3-music-platform/app/user/repository/db/model"
)

type PostDao struct {
	*gorm.DB
}

func NewPostDao(ctx context.Context) *PostDao {
	return &PostDao{NewDBClient(ctx)}
}

func (pd *PostDao) CreatePost(post *model.Post) error {
	return pd.Model(&model.Post{}).Create(&post).Error
}

func (pd *PostDao) GetPostsByAuthor(userAddr string) ([]*model.Post, error) {
	var posts []*model.Post
	log.Print("userAddr", userAddr)
	err := pd.Model(model.Post{}).Where("author_addr = ?", userAddr).Find(&posts).Error
	log.Print("GetPostsByAuthor", posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
