package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"math/big"
	"web3-music-platform/config"
	"web3-music-platform/internal/mq/messages"
	"web3-music-platform/pkg/contract"
)

func CreateMusicNFT(message []byte) error {
	var logIns = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "CreateMusicNFT",
	})
	logIns.Debugf("================== CreateMusicNFT =================")
	var req messages.CreateSongNFTReq
	var tx *types.Transaction
	var err error

	if err := json.Unmarshal(message, &req); err != nil {
		logIns.Errorf("Failed to unmarshal message: %v", err)
		return err
	}

	// Mint song NFT based on artist address
	switch req.ArtistAddr {
	case config.TestAddr1:
		tx, err = contract.SongNFTTrade.CreateMusic(contract.TransactOpts1, big.NewInt(int64(req.Amount)), big.NewInt(int64(req.Price)), req.TokenURI)
	case config.TestAddr2:
		tx, err = contract.SongNFTTrade.CreateMusic(contract.TransactOpts2, big.NewInt(int64(req.Amount)), big.NewInt(int64(req.Price)), req.TokenURI)
	default:
		err = fmt.Errorf("unsupported artist address: %s", req.ArtistAddr)
	}

	if err != nil {
		logIns.WithError(err).Error("Failed to create music NFT")
		return err
	}

	logIns.Infof("Transaction hash: %v", tx.Hash().Hex())
	return nil
}
