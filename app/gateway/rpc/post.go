package rpc

import (
	"context"
	"web3-music-platform/idl/pb"
)

func CreatePost(ctx context.Context, req *pb.PostCreateRequest) (*pb.PostCreateResponse, error) {
	res, err := UserService.CreatePost(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func GetPostsByAuthor(ctx context.Context, req *pb.GetPostsByAuthorRequest) (*pb.GetPostsByAuthorResponse, error) {
	resp, err := UserService.GetPostsByAuthor(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
