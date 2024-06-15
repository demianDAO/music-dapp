package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"web3-music-platform/config"
	"web3-music-platform/pkg/contract/sm"
)

var MPToken *sm.MPToken
var SongNFT *sm.SongNFT
var SongNFTTrade *sm.SongNFTTrade

var TransactOpts1 *bind.TransactOpts
var TransactOpts2 *bind.TransactOpts

func NewGethClient() error {
	conn, err := ethclient.Dial(config.BSC_RPC)
	if err != nil {
		return err
	}
	MPToken, err = sm.NewMPToken(common.HexToAddress(config.MPTokenAddr), conn)
	if err != nil {
		return err
	}
	SongNFT, err = sm.NewSongNFT(common.HexToAddress(config.SongNFTAddr), conn)
	if err != nil {
		return err
	}
	SongNFTTrade, err = sm.NewSongNFTTrade(common.HexToAddress(config.SongNFTTradeAddr), conn)
	if err != nil {
		return err
	}

	privateKey1, _ := crypto.HexToECDSA(config.TestPrivate1)
	privateKey2, _ := crypto.HexToECDSA(config.TestPrivate2)
	chainId, _ := conn.ChainID(context.Background())

	TransactOpts1, err = bind.NewKeyedTransactorWithChainID(privateKey1, chainId)
	if err != nil {
		return err
	}

	TransactOpts2, err = bind.NewKeyedTransactorWithChainID(privateKey2, chainId)
	if err != nil {
		return err
	}

	return nil
}
