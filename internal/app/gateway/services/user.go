package services

import (
	"context"
	"web3-music-platform/internal/app/user/models"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/utils"
)

func UserLogin(ctx context.Context, Address string) (*models.User, error) {
	res, err := UserService.UserLogin(ctx, &pb.UserLoginRequest{
		Address: Address,
	})
	if err != nil {
		return nil, err
	}
	user := utils.ToDBUser(res.User)
	return user, err
}

func UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	resp, err := UserService.UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
