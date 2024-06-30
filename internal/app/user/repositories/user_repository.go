package repositories

import (
	"context"
	"fmt"
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
	var existingUser *models.User
	// 查找是否有相同地址的用户
	if err := ud.Where("address = ?", user.Address).First(&existingUser).Error; err == nil {
		// 如果找到匹配的用户，则返回错误
		return fmt.Errorf("user with the same address already exists")
	} else if err != gorm.ErrRecordNotFound {
		// 如果查询时出现其他错误，则返回错误
		return err
	}
	// 如果没有找到匹配的用户，则创建新用户
	return ud.Create(&user).Error
}
