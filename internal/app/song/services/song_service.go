package services

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"math/big"
	"sync"
	"web3-music-platform/config"
	"web3-music-platform/internal/app/song/models"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/internal/irys"
	"web3-music-platform/internal/mq"
	"web3-music-platform/pkg/contract"
	"web3-music-platform/pkg/grpc/pb"
)

var SongServiceInstance *SongService
var SongSrvOnce sync.Once

type SongService struct {
}

func NewSongService() *SongService {
	SongSrvOnce.Do(func() {
		SongServiceInstance = &SongService{}
	})
	return SongServiceInstance
}
func (us *SongService) PurchaseSong(ctx context.Context, req *pb.PurchaseSongReq, res *pb.PurchaseSongRes) error {
	logInstance := log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "PurchaseSong",
	})
	logInstance.Infof("user = %v", req.GetUserAddr())
	logInstance.Infof("song = %v", req.GetTokenId())
	logInstance.Infof("singer = %v", req.GetSingerAddr())

	var tx *types.Transaction
	var err error

	if req.GetUserAddr() == config.TestAddr1 {
		tx, err = contract.SongNFTTrade.PurchaseSong(contract.TransactOpts1, big.NewInt(int64(req.GetTokenId())), common.HexToAddress(req.GetSingerAddr()))
	} else if req.GetUserAddr() == config.TestAddr2 {
		tx, err = contract.SongNFTTrade.PurchaseSong(contract.TransactOpts2, big.NewInt(int64(req.GetTokenId())), common.HexToAddress(req.GetSingerAddr()))
	}

	if err != nil {
		logInstance.Errorf("err = %v", err)
		return err
	}
	logInstance.Infof("tx = %v", tx)
	return nil
}

func (us *SongService) FindSongs(ctx context.Context, req *pb.FindSongsByAddrReq, res *pb.FindSongsByAddrRes) error {
	log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "FindSongs",
	}).Infof("req = %v", req)

	infosByContract, err := contract.SongNFTTrade.GetMusicInfos(&bind.CallOpts{Pending: true, From: common.HexToAddress(req.GetAddr())})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "FindSongs",
	}).Infof("infosByContract = %v", infosByContract)

	songInfos := make([]*pb.FindSongsByAddrRes_SongInfo, len(infosByContract))
	tokenIDs := make([]uint64, len(infosByContract))

	for id, info := range infosByContract {
		songInfos[id] = &pb.FindSongsByAddrRes_SongInfo{
			Id:       uint64(id),
			TokenId:  info.TokenId.Uint64(),
			TokenUri: info.TokenURI,
			Amount:   info.Balance.Uint64(),
			Price:    info.Price.Uint64(),
		}
		tokenIDs[id] = info.TokenId.Uint64()
	}

	repository := repositories.NewSongRepository(ctx)

	infosByDB, err := repository.GetSongsByTokenIDs(tokenIDs)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "FindSongs",
	}).Infof("infosByDB = %v", infosByDB)

	for id, info := range infosByDB {
		songInfos[id].Title = info.Title
		songInfos[id].ArtistAddr = info.ArtistAddr
		songInfos[id].Overview = info.Overview
	}
	res.SongInfos = songInfos
	return nil
}

func (us *SongService) DownloadSong(ctx context.Context, req *pb.DownloadSongReq, res *pb.DownloadSongRes) error {
	txId := req.GetTxId()
	bytes, err := irys.Download(irys.IrysClientInstance, ctx, txId)
	if err != nil {
		return err
	}
	res.SongBytes = bytes
	return nil
}

func (us *SongService) UploadSong(ctx context.Context, req *pb.CreateSongReq, res *pb.CreateSongRes) error {
	log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "UploadSong",
	}).Infof("req = %v", req)

	songRepo := repositories.NewSongRepository(ctx)

	tokenId, err := contract.SongNFT.CurrentID(&bind.CallOpts{Pending: true})

	var tx *types.Transaction

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "UploadSong",
	}).Infof("tokenId = %v", tokenId)

	err = songRepo.CreateSong(&models.Song{
		Title:      req.GetTitle(),
		ArtistAddr: req.GetArtistAddr(),
		Overview:   req.GetOverview(),
		TokenID:    tokenId.Uint64(),
	})

	if err != nil {
		return err
	}

	//mint song nft
	go func() {
		if req.GetArtistAddr() == config.TestAddr1 {
			tx, err = contract.SongNFTTrade.CreateMusic(contract.TransactOpts1, big.NewInt(int64(req.GetAmount())), big.NewInt(int64(req.GetPrice())), req.GetTokenUri())
		} else if req.GetArtistAddr() == config.TestAddr2 {
			tx, err = contract.SongNFTTrade.CreateMusic(contract.TransactOpts2, big.NewInt(int64(req.GetAmount())), big.NewInt(int64(req.GetPrice())), req.GetTokenUri())
		}

		if err != nil {
			log.WithFields(log.Fields{
				"pkg":  "services",
				"func": "UploadSong",
			}).Errorf("Mint err = %v", err)
		}
		log.WithFields(log.Fields{
			"pkg":  "services",
			"func": "UploadSong",
		}).Infof("tx = %v", tx)
	}()

	uploadReq, err := json.Marshal(&mq.UploadReq{
		ArtistAddr: req.GetArtistAddr(),
		TokenID:    tokenId.Uint64(),
		Data:       req.GetContent(),
	})

	err = mq.Publish(mq.RabbitMQInstance, uploadReq)

	if err != nil {
		return err
	}

	return nil
}
