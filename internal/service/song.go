package service

import (
	"context"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/repository"
)

type SongService struct {
	repo *repository.SongRepository
}

func NewSongService(_repo *repository.SongRepository) *SongService {
	return &SongService{repo: _repo}
}

func (sr SongService) CreateSong(ctx context.Context, song models.Song) error {

	err := sr.repo.CreateSong(ctx, song)
	if err != nil {
		return err
	}
	return nil
}

func (sr SongService) GetAllSongByAuthor(ctx context.Context, authorAddr string) ([]models.Song, error) {
	songs, err := sr.repo.GetAllSongsByAuthor(ctx, authorAddr)
	if err != nil {
		return nil, err
	}
	return songs, err
}
