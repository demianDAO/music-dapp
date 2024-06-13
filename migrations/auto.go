package migrations

import (
	"gorm.io/gorm"
	song_service_model "web3-music-platform/internal/app/song/models"
	user_service_model "web3-music-platform/internal/app/user/models"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(song_service_model.Song{}, user_service_model.User{}, user_service_model.Post{})
	if err != nil {
		return err
	}
	return nil
}
