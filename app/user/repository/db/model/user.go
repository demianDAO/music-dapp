package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
}

//const (
//	PassWordCost = 12 // 密码加密难度
//)
//
//// 加密密码
//func (user *User) SetPassword(password string) error {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
//	if err != nil {
//		return err
//	}
//	user.PasswordDigest = string(bytes)
//	return nil
//}
//
//// 检验密码
//func (user *User) CheckPassword(password string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
//	return err == nil
//}
