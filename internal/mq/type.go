package mq

type UploadReq struct {
	ArtistAddr string `json:"artist_addr"`
	TokenID    uint64 `json:"token_id"`
	Data       []byte `json:"data"`
}
