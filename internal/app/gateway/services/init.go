package services

import (
	"go-micro.dev/v4"
	pb2 "web3-music-platform/pkg/grpc/pb"
)

var (
	UserService pb2.UserService
	SongService pb2.SongService
)

func Init() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
	)
	songMicroService := micro.NewService(
		micro.Name("songService.client"),
	)

	UserService = pb2.NewUserService("rpcUserService", userMicroService.Client())
	SongService = pb2.NewSongService("rpcSongService", songMicroService.Client())
}
