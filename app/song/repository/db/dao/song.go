package dao

import (
	"context"
	"gorm.io/gorm"
	"log"
	"web3-music-platform/app/song/repository/db/model"
)

type SongDao struct {
	*gorm.DB
}

func NewSongDao(ctx context.Context) *SongDao {
	return &SongDao{NewDBClient(ctx)}
}

func (sd *SongDao) CreateSong(Song *model.Song) error {
	return sd.Model(&model.Song{}).Create(&Song).Error
}

func (sd SongDao) UpdateTxId(nftAddr string, tokenId uint64, txId string) error {
	log.Printf("UpdateTxId, nft:%v, tokenId：%v, txId: %v", nftAddr, tokenId, txId)
	return sd.Model(&model.Song{}).Where("nft_address = ? AND token_id = ?", nftAddr, tokenId).Update("tx_id", txId).Error
}

func (sd *SongDao) GetSongs(artistAddress string) ([]*model.Song, error) {
	var Songs []*model.Song
	log.Print("artist_address", artistAddress)
	err := sd.Model(model.Song{}).Where("artist_address = ?", artistAddress).Find(&Songs).Error
	log.Print("GetSongsByAuthor", Songs)
	if err != nil {
		return nil, err
	}
	return Songs, nil
}
