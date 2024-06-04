package rpc

import (
	"context"
	"web3-music-platform/idl/pb"
)

// UserLogin 用户登陆
func UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	res, err := UserService.UserLogin(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	resp, err := UserService.UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
