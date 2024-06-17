package handlers

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/internal/irys"
	"web3-music-platform/internal/mq/messages"
)

func SongUpload(message []byte) error {
	var logIns = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "MusicUpload",
	})
	var req messages.UploadSongReq

	err := json.Unmarshal(message, &req)

	if err != nil {
		return err
	}

	ctx := context.Background()

	songRepository := repositories.NewSongRepository(ctx)

	txId, err := irys.Upload(irys.IrysClientInstance, ctx, req.Data)

	if err != nil {
		logIns.Error(err)
		return err
	}

	logIns.Printf("tx: %v", txId)

	err = songRepository.SetTxId(req.ArtistAddr, req.TokenID, txId)

	if err != nil {
		logIns.Error(err)
		return err
	}

	return nil
}
