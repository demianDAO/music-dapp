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

var NFTContract *sm.MusicNFT
var NFTMgrContract *sm.MusicNFTMgr
var TransactOpts1 *bind.TransactOpts
var TransactOpts2 *bind.TransactOpts

func NewGethClient(rpc, nftAddr, nftMgrAddr string) error {
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	NFTContract, err = sm.NewMusicNFT(common.HexToAddress(nftAddr), conn)
	if err != nil {
		return err
	}
	NFTMgrContract, err = sm.NewMusicNFTMgr(common.HexToAddress(nftMgrAddr), conn)
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
