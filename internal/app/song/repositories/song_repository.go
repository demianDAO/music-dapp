package repositories

import (
	"context"
	"gorm.io/gorm"
	"log"
	"web3-music-platform/internal/app/song/db"
	"web3-music-platform/internal/app/song/models"
)

type SongDao struct {
	*gorm.DB
}

func NewSongDao(ctx context.Context) *SongDao {
	return &SongDao{db.NewDBClient(ctx)}
}

func (sd *SongDao) CreateSong(Song *models.Song) error {
	return sd.Model(&models.Song{}).Create(&Song).Error
}

func (sd SongDao) UpdateTxId(nftAddr string, tokenId uint64, txId string) error {
	log.Printf("UpdateTxId, nft:%v, tokenId：%v, txId: %v", nftAddr, tokenId, txId)
	return sd.Model(&models.Song{}).Where("nft_addr = ? AND token_id = ?", nftAddr, tokenId).Update("tx_id", txId).Error
}

func (sd *SongDao) GetSongs(artistAddress string) ([]*models.Song, error) {
	var Songs []*models.Song
	log.Print("artist_address", artistAddress)
	err := sd.Model(models.Song{}).Where("artist_address = ?", artistAddress).Find(&Songs).Error
	log.Print("GetSongsByAuthor", Songs)
	if err != nil {
		return nil, err
	}
	return Songs, nil
}
