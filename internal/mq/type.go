package mq

type UploadRequest struct {
	NFTAddress string `json:"nft_address"`
	TokenID    uint64 `json:"token_id"`
	Data       []byte `json:"data"`
}
