package migrations

import (
	"gorm.io/gorm"
	"web3-music-platform/app/song/repository/db/model"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(model.Song{})
	if err != nil {
		return err
	}
	return nil
}
