package rpc

import (
	"context"
	"web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"
)

func CreatePost(ctx context.Context, userAddr, content string) (*model.Post, error) {
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

func GetPosts(ctx context.Context, userAddr string) ([]*model.Post, error) {
	var posts []*model.Post
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
