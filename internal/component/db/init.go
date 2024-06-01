package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/config"
)

var DbEngine *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Music{})
	if err != nil {
		panic(err)
	}
	DbEngine = db
}

//func NewDBClient(ctx context.Context) *gorm.DB {
//	db := _db
//	return db.WithContext(ctx)
//}
