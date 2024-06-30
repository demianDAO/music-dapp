package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title      string `json:"title"`
	ArtistAddr string `json:"artist_addr"`
	Overview   string `json:"overview" gorm:"type:TEXT"`
	TokenID    uint64 `json:"token_id"`
	TxId       string `json:"tx_id"`
}
