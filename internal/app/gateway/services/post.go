package services

import (
	"context"
	"web3-music-platform/internal/app/user/models"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/utils"
)

func CreatePost(ctx context.Context, userAddr, content string) (*models.Post, error) {
	req := &pb.PostCreateRequest{
		UserAddr: userAddr,
		Content:  content,
	}
	res, err := UserService.CreatePost(ctx, req)
	postModel := res.Post
	if err != nil {
		return nil, err
	}
	return utils.ToDBPost(postModel), err
}

func GetPosts(ctx context.Context, userAddr string) ([]*models.Post, error) {
	var posts []*models.Post
	req := &pb.GetPostsRequest{
		UserAddr: userAddr,
	}
	res, err := UserService.GetPosts(ctx, req)
	if err != nil {
		return nil, err
	}
	postModels := res.GetPost()
	for _, postModel := range postModels {
		posts = append(posts, utils.ToDBPost(postModel))
	}
	return posts, err
}
