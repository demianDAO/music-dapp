package handlers

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/internal/irys"
	"web3-music-platform/internal/mq"
)

var logInstance = log.WithFields(log.Fields{
	"module": "handlers",
	"func":   "UploadAudioFile",
})

func UploadAudioFile(message []byte) error {
	var uploadReq mq.UploadReq

	err := json.Unmarshal(message, &uploadReq)

	if err != nil {
		return err
	}

	ctx := context.Background()

	songRepository := repositories.NewSongRepository(ctx)

	txId, err := irys.Upload(irys.IrysClientInstance, ctx, uploadReq.Data)

	if err != nil {
		logInstance.Error(err)
		return err
	}

	logInstance.Print("UploadAudioFile success", txId)

	err = songRepository.SetTxId(uploadReq.ArtistAddr, uploadReq.TokenID, txId)

	if err != nil {
		logInstance.Error(err)
		return err
	}

	return nil
}
