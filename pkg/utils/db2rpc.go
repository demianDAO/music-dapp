package utils

import (
	"web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/idl/pb"
)

func ToRPCUser(user *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:      uint64(user.ID),
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToDBUser(user *pb.UserModel) *model.User {
	return &model.User{
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToRPCPost(post *model.Post) *pb.PostModel {
	return &pb.PostModel{
		Id:         uint64(post.ID),
		AuthorAddr: post.AuthorAddr,
		Content:    post.Content,
	}
}

func ToDBPost(post *pb.PostModel) *model.Post {
	return &model.Post{
		AuthorAddr: post.AuthorAddr,
		Content:    post.Content,
	}
}
