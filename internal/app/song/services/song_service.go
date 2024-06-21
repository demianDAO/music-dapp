package services

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"sync"
	"web3-music-platform/internal/app/song/models"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/internal/irys"
	"web3-music-platform/internal/mq"
	"web3-music-platform/internal/mq/messages"
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

	purchaseSongNFTReq, err := json.Marshal(&messages.PurchaseSongNFTReq{
		SingerAddr: req.GetSingerAddr(),
		TokenID:    req.GetTokenId(),
		UserAddr:   req.GetUserAddr(),
	})

	err = mq.Publish(mq.RabbitMQInstance, mq.QueuesRouteKey[mq.PurchaseSongQueue], purchaseSongNFTReq)

	if err != nil {
		logInstance.Errorf("PurchaseSongQueue err = %v", err)
		return err
	}

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
	logIns := log.WithFields(log.Fields{
		"pkg":  "services",
		"func": "UploadSong",
	})
	//var sm sync.Mutex

	songRepo := repositories.NewSongRepository(ctx)

	//sm.Lock()
	tokenId, err := contract.SongNFT.CurrentID(&bind.CallOpts{Pending: true})

	if err != nil {
		logIns.Errorf("CurrentID err = %v", err)
		return err
	}
	//sm.Unlock()

	logIns.Debugf("CurrentID:%v", tokenId)

	err = songRepo.CreateSong(&models.Song{
		Title:      req.GetTitle(),
		ArtistAddr: req.GetArtistAddr(),
		Overview:   req.GetOverview(),
		TokenID:    tokenId.Uint64(),
	})

	if err != nil {
		logIns.Errorf("CreateSong err = %v", err)
		return err
	}

	go func() {
		uploadSongReq, err := json.Marshal(&messages.UploadSongReq{
			ArtistAddr: req.GetArtistAddr(),
			TokenID:    tokenId.Uint64(),
			Data:       req.GetContent(),
		})

		err = mq.Publish(mq.RabbitMQInstance, mq.QueuesRouteKey[mq.SongUploadQueue], uploadSongReq)

		if err != nil {
			logIns.Errorf("SongUpload err = %v", err)
		}
	}()

	go func() {
		createSongNFTReq, err := json.Marshal(&messages.CreateSongNFTReq{
			ArtistAddr: req.GetArtistAddr(),
			TokenURI:   req.GetTokenUri(),
			Price:      req.GetPrice(),
			Amount:     req.GetAmount(),
		})

		err = mq.Publish(mq.RabbitMQInstance, mq.QueuesRouteKey[mq.CreateSongQueue], createSongNFTReq)

		if err != nil {
			logIns.Errorf("createSongNFT err = %v", err)
		}
	}()

	return nil
}
