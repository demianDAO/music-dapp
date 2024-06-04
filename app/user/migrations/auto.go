package migrations

import (
	"gorm.io/gorm"
	"web3-music-platform/app/user/repository/db/model"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(model.User{}, model.Post{})
	if err != nil {
		return err
	}
	return nil
}
