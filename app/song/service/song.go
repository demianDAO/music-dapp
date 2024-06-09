package service

import (
	"context"
	"github.com/goccy/go-json"
	"log"
	"sync"
	"web3-music-platform/app/song/repository/db/dao"
	"web3-music-platform/app/song/repository/irys_cli"
	"web3-music-platform/app/song/repository/mq"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"
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

func (us *SongService) FindSongs(ctx context.Context, request *pb.FindSongsRequest, response *pb.FindSongsResponse) error {
	songDao := dao.NewSongDao(ctx)
	songs, err := songDao.GetSongs(request.GetArtistAddress())
	if err != nil {
		return err
	}
	var responseSongs []*pb.SongModel

	for _, song := range songs {
		responseSongs = append(responseSongs, utils.ToSongModel(song))
	}

	response.Songs = responseSongs

	return nil
}

func (us *SongService) DownloadSong(ctx context.Context, request *pb.DownloadSongRequest, response *pb.DownloadSongResponse) error {
	txId := request.GetTxId()

	bytes, err := irys_cli.Download(irys_cli.IrysClientInstance, ctx, txId)
	if err != nil {
		return err
	}
	response.SongBytes = bytes
	return nil
}

func (us *SongService) UploadSong(ctx context.Context, request *pb.CreateSongRequest, response *pb.CreateSongResponse) error {
	bytes, err := json.MarshalIndent(request.GetSong(), "", "  ")
	if err != nil {
		return err
	}
	log.Print("UploadSong", string(bytes))

	var song = utils.ToSong(request.GetSong())
	songDao := dao.NewSongDao(ctx)
	err = songDao.CreateSong(song)

	if err != nil {
		return err
	}

	uploadReq, err := json.Marshal(&mq.UploadRequest{
		NFTAddress: song.NFTAddress,
		TokenID:    song.TokenID,
		Data:       request.GetContent(),
	})

	err = mq.RabbitMQInstance.Publish(uploadReq)
	if err != nil {
		return err
	}

	return nil
}

func UploadHandler(message []byte) error {
	var uploadReq mq.UploadRequest
	err := json.Unmarshal(message, &uploadReq)
	if err != nil {
		return err
	}

	ctx := context.Background()
	songDao := dao.NewSongDao(ctx)

	txId, err := irys_cli.Upload(irys_cli.IrysClientInstance, ctx, uploadReq.Data)
	if err != nil {
		log.Print("UploadAudioFile error", err)
		return err
	}
	log.Print("UploadAudioFile success", txId)

	err = songDao.UpdateTxId(uploadReq.NFTAddress, uploadReq.TokenID, txId)
	if err != nil {
		log.Print("UpdateTxId error", err)
		return err
	}
	log.Print("UpdateTxId success")

	return nil
}
