package handlers

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"math/big"
	"web3-music-platform/config"
	"web3-music-platform/internal/mq/messages"
	"web3-music-platform/pkg/contract"
)

func PurchaseSong(message []byte) error {
	var logIns = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "PurchaseSong",
	})
	var req *messages.PurchaseSongNFTReq
	var tx *types.Transaction
	var err error

	_ = json.Unmarshal(message, &req)
	if req.UserAddr == config.TestAddr1 {
		tx, err = contract.SongNFTTrade.PurchaseSong(contract.TransactOpts1, big.NewInt(int64(req.TokenID)), common.HexToAddress(req.SingerAddr))
	} else if req.UserAddr == config.TestAddr2 {
		tx, err = contract.SongNFTTrade.PurchaseSong(contract.TransactOpts2, big.NewInt(int64(req.TokenID)), common.HexToAddress(req.SingerAddr))
	}

	if err != nil {
		logIns.WithError(err)
		return err
	}
	logIns.Infof("tx = %v", tx)

	return nil
}
