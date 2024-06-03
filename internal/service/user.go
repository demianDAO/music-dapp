package service

import (
	"context"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/repository"
)

// UserService 是用户服务接口
type UserService struct {
	repo repository.UserRepository
}

// NewUserService 创建一个新的用户服务
func NewUserService(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}

// CreateUser 创建一个新用户
func (us *UserService) CreateUser(ctx context.Context, user models.User) error {
	// 创建新用户
	err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserByAddr(ctx context.Context, address string) (*models.User, error) {
	user, err := us.repo.GetUserByAddr(ctx, address)
	if err != nil {
		return nil, err
	}
	return user, err
}
