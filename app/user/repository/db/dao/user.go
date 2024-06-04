package dao

import (
	"context"
	"log"
	"web3-music-platform/app/user/repository/db/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (ud *UserDao) FindUserByAddress(address string) (*model.User, error) {
	log.Print("FindUserByAddress", address)
	var user *model.User
	err := ud.Model(&model.User{}).
		Where("address = ?", address).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// CreateUser 创建用户
func (ud *UserDao) CreateUser(user *model.User) error {
	return ud.Model(&model.User{}).Create(&user).Error
}
