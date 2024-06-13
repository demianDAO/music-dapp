package repositories

import (
	"context"
	"log"
	"web3-music-platform/internal/app/user/db"
	"web3-music-platform/internal/app/user/models"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{db.NewDBClient(ctx)}
}

func (ud *UserDao) FindUserByAddress(address string) (*models.User, error) {
	log.Print("FindUserByAddress", address)
	var user *models.User
	err := ud.Model(&models.User{}).
		Where("address = ?", address).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

// CreateUser 创建用户
func (ud *UserDao) CreateUser(user *models.User) error {
	return ud.Model(&models.User{}).Create(&user).Error
}
