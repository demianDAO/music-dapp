package utils

import (
	user_service_model "web3-music-platform/internal/app/user/models"
	"web3-music-platform/pkg/grpc/pb"
)

func ToRPCUser(user *user_service_model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:      uint64(user.ID),
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToDBUser(user *pb.UserModel) *user_service_model.User {
	return &user_service_model.User{
		Name:    user.Name,
		Address: user.Address,
	}
}

func ToRPCPost(post *user_service_model.Post) *pb.PostModel {
	return &pb.PostModel{
		Id:       uint64(post.ID),
		UserAddr: post.UserAddr,
		Content:  post.Content,
	}
}

func ToDBPost(post *pb.PostModel) *user_service_model.Post {
	return &user_service_model.Post{
		UserAddr: post.UserAddr,
		Content:  post.Content,
	}
}
