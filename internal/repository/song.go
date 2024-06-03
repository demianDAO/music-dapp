package repository

import (
	"context"
	"gorm.io/gorm"
	"web3-music-platform/internal/api/models"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (sr SongRepository) GetAllSongsByAuthor(ctx context.Context, authorAddr string) ([]models.Song, error) {
	var songs []models.Song

	err := sr.db.WithContext(ctx).Where("author_address = ?", authorAddr).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, err
}

func (sr SongRepository) CreateSong(ctx context.Context, song models.Song) error {
	err := sr.db.WithContext(ctx).Create(song).Error
	if err != nil {
		return err
	}
	return nil
}
