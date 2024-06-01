package repository

import (
	"context"
	"gorm.io/gorm"
	"web3-music-platform/internal/api/models"
)

// UserRepository 是用户仓库接口
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建一个新的用户仓库
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// CreateUser 创建一个新用户
func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	err := ur.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID 根据用户ID获取用户信息
func (ur *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := ur.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
