package service

import (
	"context"
	"github.com/goccy/go-json"
	"log"
	"sync"
	"web3-music-platform/app/song/repository/db/dao"
	"web3-music-platform/app/song/repository/db/decentralised_storage"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"
)

var SongServiceInstance *SongService
var SongSrvOnce sync.Once

type SongService struct {
}

func (us *SongService) UploadSong(ctx context.Context, request *pb.CreateSongRequest, response *pb.CreateSongResponse) error {
	bytes, err := json.MarshalIndent(request.GetSong(), "", "  ")
	if err != nil {
		return err
	}
	log.Print("UploadSong", string(bytes))

	var song = utils.ToSong(request.GetSong())
	//upload audio data
	txId, err := decentralised_storage.IrysClientInstance.UploadAudioFile(ctx, request.GetContent())
	if err != nil {
		log.Print("UploadAudioFile error", err)
		return err
	}
	song.IrysTxId = txId

	err = dao.NewSongDao(ctx).CreateSong(song)
	if err != nil {
		return err
	}
	response.Song = utils.ToSongModel(song)
	return nil
}

func NewSongService() *SongService {
	SongSrvOnce.Do(func() {
		SongServiceInstance = &SongService{}
	})
	return SongServiceInstance
}
