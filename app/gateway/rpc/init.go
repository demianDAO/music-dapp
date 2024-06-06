package rpc

import (
	"go-micro.dev/v4"
	"web3-music-platform/idl/pb"
)

var (
	UserService pb.UserService
	SongService pb.SongService
)

func InitRPC() {
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
	)
	songMicroService := micro.NewService(
		micro.Name("songService.client"),
	)

	UserService = pb.NewUserService("rpcUserService", userMicroService.Client())
	SongService = pb.NewSongService("rpcSongService", songMicroService.Client())
}
