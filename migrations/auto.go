package migrations

import (
	"gorm.io/gorm"
	song_servicc "web3-music-platform/app/song/repository/db/model"
	user_servicc "web3-music-platform/app/user/repository/db/model"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(song_servicc.Song{}, user_servicc.User{}, user_servicc.Post{})
	if err != nil {
		return err
	}
	return nil
}
