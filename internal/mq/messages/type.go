package messages

type UploadSongReq struct {
	ArtistAddr string `json:"artist_addr"`
	TokenID    uint64 `json:"token_id"`
	Data       []byte `json:"data"`
}

type CreateSongNFTReq struct {
	ArtistAddr string `json:"artist_addr"`
	TokenURI   string `json:"token_uri"`
	Amount     uint64 `json:"amount"`
	Price      uint64 `json:"price"`
}

type PurchaseSongNFTReq struct {
	SingerAddr string `json:"singer_addr"`
	TokenID    uint64 `json:"token_id"`
	UserAddr   string `json:"user_addr"`
}
