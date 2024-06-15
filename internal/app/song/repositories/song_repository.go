package repositories

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"web3-music-platform/internal/app/song/db"
	"web3-music-platform/internal/app/song/models"
)

type SongRepository struct {
	*gorm.DB
}

func NewSongRepository(ctx context.Context) *SongRepository {
	return &SongRepository{db.NewDBClient(ctx)}
}

func (sr *SongRepository) CreateSong(song *models.Song) error {
	return sr.Model(&models.Song{}).Create(&song).Error
}

func (sr *SongRepository) SetTxId(artistAddr string, tokenId uint64, txId string) error {
	log.WithFields(log.Fields{
		"pkg":  "repositories",
		"func": "UpdateTxId",
	}).Infof("tokenId = %v,txId = %v", tokenId, txId)

	return sr.Model(&models.Song{}).Where("token_id = ? AND artist_addr = ?", tokenId, artistAddr).Update("tx_id", txId).Error
}

func (sr *SongRepository) GetSongsByAddr(artistAddr string) ([]*models.Song, error) {

	log.WithFields(log.Fields{
		"pkg":  "repositories",
		"func": "UpdateTxId",
	}).Infof("artistAddr = %v", artistAddr)

	var songs []*models.Song

	err := sr.Model(&models.Song{}).Where("artist_addr = ?", artistAddr).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (sr *SongRepository) GetSongsByTokenIDs(tokenIDs []uint64) ([]*models.Song, error) {
	log.WithFields(log.Fields{
		"pkg":  "repositories",
		"func": "GetSongsByTokenIDs",
	}).Infof("tokenIDs = %v", tokenIDs)

	var songs []*models.Song
	err := sr.Model(&models.Song{}).Where("token_id IN (?)", tokenIDs).Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}