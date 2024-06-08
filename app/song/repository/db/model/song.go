package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title         string `json:"title"`
	ArtistAddress string `json:"artist_address"`
	Overview      string `json:"overview"`
	NFTAddress    string `json:"nft_address"`
	TokenID       uint64 `json:"token_id"`
	TxId          string `json:"tx_id"`
}
