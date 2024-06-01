package service

import (
	"context"
	"errors"
	"strconv"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/repository"
)

// UserService 是用户服务接口
type UserService struct {
	userRepository *repository.UserRepository
}

// NewUserService 创建一个新的用户服务
func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

// CreateUser 创建一个新用户
func (us *UserService) CreateUser(ctx context.Context, user *models.User) error {
	// 检查用户是否已存在
	existingUser, err := us.userRepository.GetUserByID(ctx, strconv.Itoa(int(user.ID)))
	if err == nil && existingUser != nil {
		return errors.New("用户已存在")
	}

	// 创建新用户
	err = us.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
